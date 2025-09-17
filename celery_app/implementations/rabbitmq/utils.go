package rabbit

import (
	q "celery_client/celery_app/implementations/rabbitmq/queue"
	"fmt"
)

//func (b *RabbitMQ) newCeleryTask(rawTask amqp.Delivery) Task {
//	//fmt.Println(rawTask)
//
//	return Task{
//		tmp:    rawTask,
//		Header: ,
//		Body:   protocol.Task{},
//	}
//}

func (b *RabbitMQ) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
}

func (b *RabbitMQ) declareQueue(queue q.Queue) {
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
