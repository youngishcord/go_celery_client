package tasks

import (
	"context"
	"errors"
	"go_celery_client/celery/internal/task"
	"go_celery_client/celery/protocol"
	"time"
)

type AddTask struct {
	a, b float64
}

func NewAddTask(rawTask *protocol.CeleryTask) (task.Task, error) {
	t := AddTask{}

	args := rawTask.Body.Args

	if a, ok := args[0].(float64); !ok {
		return nil, errors.New("INVALID_ARGUMENT")
	} else {
		t.a = a
	}

	if b, ok := args[1].(float64); !ok {
		return nil, errors.New("INVALID_ARGUMENT")
	} else {
		t.b = b
	}

	return &t, nil
}

func (t *AddTask) Run(ctx context.Context) (any, error) {
	time.Sleep(7 * time.Second)
	return t.a + t.b, nil
}

func (t *AddTask) Name() string {
	return "add"
}

// TODO: Fix name on celery doc. Delay only for sync task apply
func (t *AddTask) Delay() error {
	return nil
}
