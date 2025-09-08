package broker

import amqp "github.com/rabbitmq/amqp091-go"

type Broker interface {
	Connect(queues []string) error
	TaskChannel() chan amqp.Delivery // Я только что понял, что этот интерфейс не
	// будет работать с Redis, поскольку он не универсален
	Connection() amqp.Connection
}
