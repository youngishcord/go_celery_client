package redis

import (
	interf "celery_client/celery_app/core/interfaces"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Task struct {
}

func (t *Task) Reject() {
	panic("implement me")
}

func (t *Task) Nack() {
	panic("implement me")
}

func (t *Task) Ack() {
	panic("implement me")
}

func (t *Task) Name() string {
	panic("implement me")
}

func NewTask(rawTask amqp.Delivery) interf.Tasks {
	panic("implement me")
}
