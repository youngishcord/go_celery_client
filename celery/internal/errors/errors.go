package errors

import "errors"

// celery errors
var (
	ErrNotRegistered         = errors.New("NotRegistered")
	ErrFailOnRunningTask     = errors.New("FailOnRunningTask")
	ErrSoftTimeLimitExceeded = errors.New("SoftTimeLimitExceeded")
	ErrHardTimeLimitExceeded = errors.New("HardTimeLimitExceeded")
)
