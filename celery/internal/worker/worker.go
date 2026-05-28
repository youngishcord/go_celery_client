package worker

import (
	"context"
	"fmt"
	"go_celery_client/celery/internal/task"
	"go_celery_client/celery/protocol"
	"log"
)

type App interface {
	Consume() <-chan *protocol.CeleryTask
	PublishResult(ctx context.Context, result any, celeryTask *protocol.CeleryTask) error
	MakeTask(ctx context.Context, task *protocol.CeleryTask) (task.Task, error)
}

type CeleryWorker struct {
	index int
	app   App

	closeCh <-chan struct{}
}

func NewCeleryWorker(index int, app App, closeCh <-chan struct{}) (*CeleryWorker, error) {
	return &CeleryWorker{
		index:   index,
		app:     app,
		closeCh: closeCh,
	}, nil
}

func (w *CeleryWorker) Start() error {
	log.Println("Starting Celery worker ", w.index)
	for {
		select {
		case t, ok := <-w.app.Consume():
			log.Println("worker ", w.index, " receive task: ", ok)
			log.Println(t)
			err := w.processTask(t)
			if err != nil {
				log.Println(err)
			}
		case <-w.closeCh:
			log.Println(fmt.Sprintf("Worker %d stopped", w.index))
			return nil
		}
	}
}
