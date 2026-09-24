package exceptions

import (
	"errors"
	"fmt"
	"sync"

	e "go_celery_client/celery/internal/errors"
)

type ExceptionRegistry struct {
	mu      sync.RWMutex
	storage map[error]BaseException
}

func NewExceptionRegistry() *ExceptionRegistry {
	r := &ExceptionRegistry{
		storage: make(map[error]BaseException),
	}

	r.RegisterException(e.ErrNotRegistered, BaseException{
		ExceptionType:   e.ErrNotRegistered.Error(),
		ExceptionModule: module,
	})
	r.RegisterException(e.ErrFailOnRunningTask, BaseException{
		ExceptionType:   e.ErrFailOnRunningTask.Error(),
		ExceptionModule: module,
	})
	r.RegisterException(e.ErrSoftTimeLimitExceeded, BaseException{
		ExceptionType:   e.ErrSoftTimeLimitExceeded.Error(),
		ExceptionModule: module,
	})
	r.RegisterException(e.ErrHardTimeLimitExceeded, BaseException{
		ExceptionType:   e.ErrHardTimeLimitExceeded.Error(),
		ExceptionModule: module,
	})

	return r
}

// RegisterException регистрирует соответствие sentinel-ошибки Go Python-исключению.
// Поиск выполняется через errors.Is, поэтому работают и обернутые ошибки (%w).
func (r *ExceptionRegistry) RegisterException(err error, exc BaseException) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.storage[err]; ok {
		return fmt.Errorf("%q already registered", err.Error())
	}

	r.storage[err] = exc
	return nil
}

func (r *ExceptionRegistry) find(err error) (BaseException, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if exc, ok := r.storage[err]; ok {
		return exc, true
	}

	for sentinel, exc := range r.storage {
		if errors.Is(err, sentinel) {
			return exc, true
		}
	}

	return BaseException{}, false
}

// ExceptionInfo преобразует ошибку в celery-совместимое исключение.
// Если message пустой, сообщением становится текст ошибки.
func (r *ExceptionRegistry) ExceptionInfo(err error, message []string, args []any, kwargs map[string]any) *ExceptionInfo {
	exc := Exception
	if found, ok := r.find(err); ok {
		exc = found
	}
	if message == nil {
		message = []string{err.Error()}
	}
	return NewExceptionInfo(exc.ExceptionType, message, exc.ExceptionModule, args, kwargs)
}
