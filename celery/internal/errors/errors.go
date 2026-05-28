package errors

import "errors"

// celery errors
var (
	ErrNotRegistered         = errors.New("NotRegistered")
	ErrSoftTimeLimitExceeded = errors.New("SoftTimeLimitExceeded")
	ErrHardTimeLimitExceeded = errors.New("HardTimeLimitExceeded")
)
