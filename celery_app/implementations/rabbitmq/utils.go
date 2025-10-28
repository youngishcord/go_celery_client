package rabbit

import (
	q "celery_client/celery_app/implementations/rabbitmq/queue"
	"fmt"

	protocol "celery_client/celery_app/core/dto/protocol"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewTask(rawTask amqp.Delivery) protocol.CeleryTask {

	body, err := protocol.ParsePayload(rawTask.Body)
	if err != nil {
		panic(err)
	}

	header, err := protocol.ParseHeader(rawTask.Headers)
	if err != nil {
		panic(err)
	}

	correlationID, err := uuid.Parse(rawTask.CorrelationId)
	replyTo, err := uuid.Parse(rawTask.ReplyTo)

	return protocol.CeleryTask{
		ContentEncoding: rawTask.ContentEncoding,
		ContentType:     rawTask.ContentType,
		Body:            body,
		Headers:         header,
		Properties: protocol.Properties{
			CorrelationID: correlationID,
			DeliveryTag:   rawTask.DeliveryTag,
			ReplyTo:       replyTo,
			DeliveryMode:  rawTask.DeliveryMode,
			DeliveryInfo: protocol.DeliveryInfo{
				Exchange:   rawTask.Exchange,
				RoutingKey: rawTask.RoutingKey,
			},
			Priority:     rawTask.Priority,
			BodyEncoding: rawTask.ContentEncoding,
		},
	}
}

func (b *Rabbit) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
}

func (b *Rabbit) declareQueue(queue q.Queue) {
	if b.Consumer == nil {
		panic("CHANNEL NOT OPEN")
	}

	_, err := b.Consumer.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	)
	if err != nil {
		panic("QUEUE WAS NOT DECLARED")
	}

}
