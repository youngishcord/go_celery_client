package exceptions

import (
	"errors"
	e "go_celery_client/celery/internal/errors"
)

const (
	module = "celery.exceptions"
)

var Exception BaseException = BaseException{
	ExceptionType:   "Exception",
	ExceptionModule: "builtins",
}

type ExceptionRegistry struct {
	storage map[string]BaseException
}

func NewExceptionRegistry() *ExceptionRegistry {
	return &ExceptionRegistry{
		storage: map[string]BaseException{
			e.ErrNotRegistered.Error(): {
				ExceptionType:   e.ErrNotRegistered.Error(),
				ExceptionModule: module,
			},
			e.ErrSoftTimeLimitExceeded.Error(): {
				ExceptionType:   e.ErrSoftTimeLimitExceeded.Error(),
				ExceptionModule: module,
			},
		},
	}
}

func (r ExceptionRegistry) RegisterNewExceptions(ex map[string]BaseException) error {
	for key, value := range ex {
		if _, ok := r.storage[key]; ok {
			return errors.New(key + " already exists")
		}
		r.storage[key] = value
	}
	return nil
}

// GetException make celery-like exception from Go error
func (r ExceptionRegistry) GetException(err error, exceptionMessage []string, args []any, kwargs map[string]any) *ExceptionInfo {
	var exception BaseException = Exception
	if exc, ok := r.storage[err.Error()]; ok {
		exception = exc
	}
	return NewExceptionInfo(
		exception.ExceptionType,
		exceptionMessage,
		exception.ExceptionModule,
		args,
		kwargs,
	)
}
