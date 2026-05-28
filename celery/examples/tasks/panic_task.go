package tasks

import (
	"context"
	"go_celery_client/celery/internal/task"
	"go_celery_client/celery/protocol"
)

type PanicTask struct{}

func NewPanicTask(rawTask *protocol.CeleryTask) (task.Task, error) {
	return &PanicTask{}, nil
}

func (t *PanicTask) Run(ctx context.Context) (any, error) {
	panic("TEST PANIC IN TASK")

	return nil, nil
}

func (t *PanicTask) Name() string {
	return "panic_task"
}

func (t *PanicTask) Delay() error {
	return nil
}
