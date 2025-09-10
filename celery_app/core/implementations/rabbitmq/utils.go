package rabbit

import (
	q "celery_client/celery_app/core/implementations/rabbitmq/queue"
	"celery_client/celery_app/dto"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (b *RabbitMQBroker) newCeleryTask(rawTask amqp.Delivery) dto.CeleryRawTask {
	fmt.Println(rawTask)

	return dto.CeleryRawTask{
		TestMessage: "THIS IS TEST MESSAGE",
	}
}

func (b *RabbitMQBroker) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
}

func (b *RabbitMQBroker) declareQueue(queue q.Queue) {
	if b.Channel == nil {
		panic("CHANNEL NOT OPEN")
	}

	_, err := b.Channel.QueueDeclare(
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
