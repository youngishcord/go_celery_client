package tasks

import (
	"context"
	"fmt"
	"go_celery_client/celery/protocol"
	"go_celery_client/celery/task"
)

type ErrorTask struct{}

func NewErrorTask(rawTask *protocol.CeleryTask) (task.Task, error) {
	return &ErrorTask{}, nil
}

func (t *ErrorTask) Run(ctx context.Context) (any, error) {
	return nil, fmt.Errorf("ErrorTask Run")
}

func (t *ErrorTask) Name() string {
	return "error"
}
