package worker

import (
	"context"
	e "go_celery_client/celery/internal/errors"
	"go_celery_client/celery/protocol"
	"log"
)

func (w *CeleryWorker) processTask(celeryTask *protocol.CeleryTask) error {
	hardCtx, cancel := MakeContext(context.Background(), celeryTask.Headers.TimeLimit.Hard)
	defer cancel()

	softCtx, softCancel := MakeContext(hardCtx, celeryTask.Headers.TimeLimit.Soft)
	defer softCancel()

	task, err := w.app.MakeTask(hardCtx, celeryTask)
	if err != nil {
		log.Println("TASK NOT FOUND")
		//err = w.app..PublishException(
		//	context.Background(),
		//	exceptions.GetException(e.ErrNotRegistered,
		//		[]string{celeryTask.Headers.Task}),
		//	celeryTask,
		//	"test trace",
		//)
		//if err != nil {
		//	log.Println(err)
		//}
		return err
	}

	taskResult, err := RunWithTimeout(softCtx, hardCtx, task.Run)
	if err != nil {
		//err := a.Backend.PublishException(
		//	context.Background(), // Что тут делать с контекстом?
		//	exceptions.GetException(err, []string{}),
		//	celeryTask,
		//	"test trace",
		//)
		//if err != nil {
		//	log.Println(err)
		//}
		return err
	}

	err = w.app.PublishResult(softCtx, taskResult, celeryTask)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func RunWithTimeout(softCtx context.Context, hardCtx context.Context, fn func(ctx2 context.Context) (any, error)) (any, error) {
	done := make(chan any)
	errCh := make(chan error)

	go func() {
		res, err := fn(softCtx)
		if err != nil {
			errCh <- err
		}
		done <- res
	}()

	select {
	case res := <-done:
		return res, nil
	case err := <-errCh:
		return nil, err
	case <-softCtx.Done():
		return nil, e.ErrSoftTimeLimitExceeded
	case <-hardCtx.Done():
		return nil, e.ErrSoftTimeLimitExceeded
	}
}
