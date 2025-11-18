package celery_app

import (
	e "celery_client/celery_app/core/errors"
	"context"
	"time"
)

func MakeContext(ctx context.Context, timeout *time.Duration) (context.Context, context.CancelFunc) {
	if timeout != nil && *timeout > 0 {
		return context.WithTimeout(ctx, *timeout)
	}
	return context.WithCancel(ctx)
}

// TODO: наверное надо сделать отдельно сущность воркера и вынести в него этот метод
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
