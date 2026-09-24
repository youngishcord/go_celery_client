package tasks

import (
	"context"
	"errors"
	"fmt"
	"go_celery_client/celery/protocol"
	"go_celery_client/celery/task"
)

// ErrCustom — пример пользовательской ошибки, которую можно зарегистрировать
// в реестре исключений приложения и возвращать из задачи.
var ErrCustom = errors.New("custom error")

type ErrorTask struct{}

func NewErrorTask(rawTask *protocol.CeleryTask) (task.Task, error) {
	return &ErrorTask{}, nil
}

func (t *ErrorTask) Run(ctx context.Context) (any, error) {
	return nil, fmt.Errorf("ErrorTask Run: %w", ErrCustom)
}

func (t *ErrorTask) Name() string {
	return "error_task"
}
