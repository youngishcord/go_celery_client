package broker

import amqp "github.com/rabbitmq/amqp091-go"

type Broker interface {
	Connect(queues []string) error
	TaskChannel() chan amqp.Delivery
}
