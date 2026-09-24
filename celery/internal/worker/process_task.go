package worker

import (
	"context"
	"fmt"
	"go_celery_client/celery/internal/errors"
	"go_celery_client/celery/protocol"
	"log"
	"runtime/debug"
)

func (w *CeleryWorker) processTask(celeryTask *protocol.CeleryTask) error {
	hardCtx, cancel := MakeContext(context.Background(), celeryTask.Headers.TimeLimit.Hard)
	defer cancel()

	softCtx, softCancel := MakeContext(hardCtx, celeryTask.Headers.TimeLimit.Soft)
	defer softCancel()

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic: %v\nstack: %s", r, debug.Stack())
			log.Println(err)
			e := w.app.PublishException(
				context.Background(),
				w.app.ExceptionInfo(err, nil, nil, nil),
				celeryTask,
				"",
			)
			if e != nil {
				log.Println(e)
			}
		}
	}()

	task, err := w.app.MakeTask(hardCtx, celeryTask)
	if err != nil {
		message := []string{err.Error()}
		if err == errors.ErrNotRegistered {
			message = []string{celeryTask.Headers.Task}
		}

		e := w.app.PublishException(
			softCtx,
			w.app.ExceptionInfo(err, message, nil, nil),
			celeryTask,
			"",
		)
		if e != nil {
			return e
		}
		return err
	}

	taskResult, err := RunWithTimeout(softCtx, hardCtx, task.Run)
	if err != nil {
		e := w.app.PublishException(
			softCtx,
			w.app.ExceptionInfo(err, nil, nil, nil),
			celeryTask,
			"",
		)
		if e != nil {
			return e
		}
		return err
	}

	err = w.app.PublishResult(softCtx, taskResult, celeryTask)
	if err != nil {
		e := w.app.PublishException(softCtx, "", celeryTask, "")
		if e != nil {
			return e
		}
		return err
	}

	return nil
}

func RunWithTimeout(softCtx context.Context, hardCtx context.Context, fn func(ctx2 context.Context) (any, error)) (any, error) {
	done := make(chan any)
	errCh := make(chan error)
	defer close(done)
	defer close(errCh)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errCh <- fmt.Errorf("panic: %v\nstack: %s", r, debug.Stack())
			}
		}()
		res, err := fn(softCtx)
		if err != nil {
			errCh <- err
			return
		}
		done <- res
	}()

	select {
	case res := <-done:
		return res, nil
	case err := <-errCh:
		return nil, err
		//case <-softCtx.Done():
		//	return nil, e.ErrSoftTimeLimitExceeded
		//case <-hardCtx.Done():
		//	return nil, e.ErrHardTimeLimitExceeded
	}
}
