package backend

import amqp "github.com/rabbitmq/amqp091-go"

// AMQPBackend Переиспользование подключения из broker
type AMQPBackend struct {
	AmqpChan *amqp.Channel
}

func NewAMQPBackend(amqpChan *amqp.Channel) *AMQPBackend {
	return &AMQPBackend{
		AmqpChan: amqpChan,
	}
}

func (b *AMQPBackend) Connect() error {
	// TODO: Скорее всего пустая функция?
	//  Надо подумать над подключением к редису
	panic("implement me")
}

func (b *AMQPBackend) MakeResult() error {

	return nil
}
