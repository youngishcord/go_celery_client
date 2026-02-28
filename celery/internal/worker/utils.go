package worker

import (
	"context"
	"time"
)

func MakeContext(ctx context.Context, timeout *time.Duration) (context.Context, context.CancelFunc) {
	if timeout != nil && *timeout > 0 {
		return context.WithTimeout(ctx, *timeout)
	}
	return context.WithCancel(ctx)
}
