package tasks

import (
	"context"
	"go_celery_client/celery/internal/task"
	"go_celery_client/celery/protocol"
	"log"
)

type InfTask struct{}

func NewInfTask(rawTask *protocol.CeleryTask) (task.Task, error) {
	return &InfTask{}, nil
}

func (t *InfTask) Run(ctx context.Context) (any, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			log.Println("inf_task")
		}
	}
	return nil, nil
}

func (t *InfTask) Name() string {
	return "inf_task"
}
