package rabbit

import (
	protocol "celery_client/celery_app/core/dto/protocol"
	"celery_client/celery_app/implementations/rabbitmq/router"
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (b *Rabbit) ConsumeTask() <-chan protocol.CeleryTask {
	return b.TaskCh
}

// Тут обработка отправки по цепочке, в дальнейшем, возможно, нужно будет отделить реализацию от других кейсов.
func (b *Rabbit) PublishTask(ctx context.Context, celeryTask protocol.CeleryTask) error {

	rawBody, err := json.Marshal(celeryTask.Body)
	if err != nil {
		return err
	}

	headers, err := celeryTask.Headers.MakeMap()
	if err != nil {
		return err
	}

	// nextTask := celeryTask.Body.Emb.Chain[len(celeryTask.Body.Emb.Chain)-1]

	err = b.declareQueue(router.Queue{
		Name:    celeryTask.Properties.DeliveryInfo.RoutingKey,
		Durable: true,
	})
	if err != nil {
		return err
	}

	err = b.Publisher.Publish(
		// ctx,
		celeryTask.Properties.DeliveryInfo.Exchange,
		celeryTask.Properties.DeliveryInfo.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:     celeryTask.ContentType,
			CorrelationId:   celeryTask.Properties.CorrelationID.String(),
			Body:            rawBody,
			Headers:         amqp.Table(headers),
			ContentEncoding: celeryTask.ContentEncoding,
			DeliveryMode:    celeryTask.Properties.DeliveryMode,
			Priority:        celeryTask.Properties.Priority,
			ReplyTo:         celeryTask.Properties.ReplyTo.String(),
			// ReplyTo:         nextTask.Opt.ReplyTo,
			Timestamp: time.Now().UTC(),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Ack basic acknowledge function for celery task
func (b *Rabbit) Ack(task protocol.CeleryTask) error {
	err := b.Consumer.Ack(*task.Properties.DeliveryTag, false)
	if err != nil {
		return err
	}
	return nil
}

// Reject basic reject function for celery task
func (b *Rabbit) Reject(task protocol.CeleryTask, requeue bool) error {
	err := b.Consumer.Reject(*task.Properties.DeliveryTag, requeue)
	if err != nil {
		return err
	}
	return nil
}

// Nack negatively acknowledges a delivery by its delivery tag.
func (b *Rabbit) Nack(task protocol.CeleryTask, requeue bool) error {
	err := b.Consumer.Nack(*task.Properties.DeliveryTag, false, requeue)
	if err != nil {
		return err
	}
	return nil
}
