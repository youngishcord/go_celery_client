package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"go_celery_client/celery/protocol"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (a *RabbitAdapter) PublishResult(ctx context.Context, result any, celeryTask *protocol.CeleryTask) error {
	body, err := json.Marshal(protocol.NewCeleryResult(protocol.SUCCESS, result, "", celeryTask.Headers.Id))
	if err != nil {
		return err
	}

	err = a.PublishCh.PublishWithContext(
		ctx,
		"",
		celeryTask.Properties.ReplyTo.String(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: celeryTask.Properties.CorrelationID.String(),
			Timestamp:     time.Now().UTC(),
			Body:          body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *RabbitAdapter) PublishException(ctx context.Context, result any, celeryTask *protocol.CeleryTask, trace string) error {
	body, err := json.Marshal(protocol.NewCeleryResult(protocol.FAILURE, result, trace, celeryTask.Headers.Id))
	if err != nil {
		return err
	}

	err = a.PublishCh.PublishWithContext(
		ctx,
		"",
		celeryTask.Properties.ReplyTo.String(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: celeryTask.Properties.CorrelationID.String(),
			Body:          body,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
