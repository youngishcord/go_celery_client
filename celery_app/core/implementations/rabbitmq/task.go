package rabbit

import (
	protocol "celery_client/celery_app/core/implementations/rabbitmq/protocol"
	interf "celery_client/celery_app/core/interfaces"

	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: тут надо понять, какие именно параметры нужно сохранять
type Task struct {
	tmp amqp.Delivery

	Header protocol.Header
	Body   protocol.Task
}

func (t *Task) Reject() {
	err := t.tmp.Reject(false)
	if err != nil {
		panic(err)
	}
}

func (t *Task) Nack() {
	err := t.tmp.Nack(false, true)
	if err != nil {
		panic(err)
	}
}

func (t *Task) Ack() {
	err := t.tmp.Ack(false)
	if err != nil {
		panic(err)
	}
}

func (t *Task) Name() string {
	return t.Header.Task
}

func NewTask(rawTask amqp.Delivery) interf.Tasks {
	body, err := protocol.ParseTask(rawTask.Body)
	if err != nil {
		panic(err)
	}
	return &Task{
		tmp:    rawTask,
		Header: protocol.ParseHeader(rawTask.Headers),
		Body:   body,
	}
}
