package adapter

import (
	"context"
	"fmt"
	"go_celery_client/celery/protocol"
)

func (a *RabbitAdapter) PublishResult(ctx context.Context, result any, celeryTask *protocol.CeleryTask) error {
	//fmt.Println("Rabbit adapter publish result")
	fmt.Println("implement me")
	return nil
}

func (a *RabbitAdapter) PublishException() error {
	return nil
}
