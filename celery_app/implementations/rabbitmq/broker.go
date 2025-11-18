package rabbit

import (
	protocol "celery_client/celery_app/core/dto/protocol"
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (b *Rabbit) ConsumeTask() <-chan protocol.CeleryTask {
	return b.TaskCh
}

func (b *Rabbit) PublishTask(ctx context.Context, celeryTask protocol.CeleryTask) error {

	rawBody, err := json.Marshal(celeryTask.Body)

	headers, err := json.Marshal(celeryTask.Headers)

	b.Publisher.PublishWithContext(
		ctx,
		"",
		celeryTask.Properties.ReplyTo.String(),
		false,
		false,
		amqp.Publishing{
			ContentType:     celeryTask.ContentType,
			CorrelationId:   celeryTask.Properties.CorrelationID.String(),
			Body:            rawBody,
			Headers:         amqp.Table{},
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// Ack basic acknowledge function for celery task
func (b *Rabbit) Ack(task protocol.CeleryTask) error {
	err := b.Consumer.Ack(task.Properties.DeliveryTag, false)
	if err != nil {
		return err
	}
	return nil
}

// Reject basic reject function for celery task
func (b *Rabbit) Reject(task protocol.CeleryTask, requeue bool) error {
	err := b.Consumer.Reject(task.Properties.DeliveryTag, requeue)
	if err != nil {
		return err
	}
	return nil
}

// Nack negatively acknowledges a delivery by its delivery tag.
func (b *Rabbit) Nack(task protocol.CeleryTask, requeue bool) error {
	err := b.Consumer.Nack(task.Properties.DeliveryTag, false, requeue)
	if err != nil {
		return err
	}
	return nil
}
