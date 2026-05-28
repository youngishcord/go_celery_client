package app

import (
	"context"
	e "go_celery_client/celery/internal/errors"
	"go_celery_client/celery/protocol"
	"go_celery_client/celery/task"
)

// MakeTask получает на вход параметры, находит конструктор задачи, формирует
// структуру и возвращает ее для дальнейшей обработки.
func (a *CeleryApp) MakeTask(ctx context.Context, task *protocol.CeleryTask) (task.Task, error) {
	constructor, ok := a.taskRegistry[task.Headers.Task]
	if !ok {
		return nil, e.ErrNotRegistered
	}
	return constructor(task)
}
