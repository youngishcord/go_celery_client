package adapter

import (
	protocol "go_celery_client/celery/protocol"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// NewTask create a Celery task from a raw amqp protocol
func NewTask(rawTask amqp.Delivery) (*protocol.CeleryTask, error) {
	body, err := protocol.ParsePayload(rawTask.Body)
	if err != nil {
		return nil, err
	}

	header, err := protocol.ParseHeader(rawTask.Headers)
	if err != nil {
		return nil, err
	}

	correlationID, err := uuid.Parse(rawTask.CorrelationId)
	replyTo, err := uuid.Parse(rawTask.ReplyTo)

	return &protocol.CeleryTask{
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
	}, nil
}

//func (b *RabbitAdapter) url() string {
//	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
//}

//func (a *RabbitAdapter) declareQueue(queue q.Queue) {
//	if a.Client == nil {
//		panic("CHANNEL NOT OPEN")
//	}
//
//	_, err := a.Client.QueueDeclare(
//		queue.Name,
//		queue.Durable,
//		queue.AutoDelete,
//		queue.Exclusive,
//		queue.NoWait,
//		queue.Args,
//	)
//	if err != nil {
//		panic("QUEUE WAS NOT DECLARED")
//	}
//
//}
