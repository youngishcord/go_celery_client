package app

import (
	"go_celery_client/celery/exceptions"
)

// RegisterException регистрирует соответствие Go-ошибки Python-исключению,
// которое будет возвращено как результат задачи при ошибке.
func (a *CeleryApp) RegisterException(err error, exc exceptions.BaseException) error {
	return a.exceptionRegistry.RegisterException(err, exc)
}

// ExceptionInfo преобразует Go-ошибку в celery-совместимое исключение.
func (a *CeleryApp) ExceptionInfo(err error, message []string, args []any, kwargs map[string]any) *exceptions.ExceptionInfo {
	return a.exceptionRegistry.ExceptionInfo(err, message, args, kwargs)
}
