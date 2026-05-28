package app

import (
	"context"
	"go_celery_client/celery/protocol"
)

func (a *CeleryApp) Consume() <-chan *protocol.CeleryTask {
	return a.broker.Consume()
}

func (a *CeleryApp) PublishResult(ctx context.Context, result any, celeryTask *protocol.CeleryTask) error {

	err := a.broker.Ack(celeryTask)
	if err != nil {
		return err
	}

	return a.backend.PublishResult(ctx, result, celeryTask)
}
