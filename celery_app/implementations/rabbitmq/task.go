package rabbit

import (
	amqp_protocol "celery_client/celery_app/implementations/rabbitmq/protocol"
	baseTask "celery_client/celery_app/task"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: тут надо понять, какие именно параметры нужно сохранять, а что отбросить
//type Task struct {
//	tmp amqp.Delivery // Мне не нравится, что тут лежит полноценная структура.
//	// Надо взять из нее все что надо и убрать отсюда
//
//	Header amqp_protocol.Header
//	Body   amqp_protocol.Task
//}
//
//func (t *Task) CorrelationID() string {
//	return t.tmp.CorrelationId
//}
//
//// FIXME: Тут надо убрать темповую переменную
//func (t *Task) ReplyTo() string {
//	return t.tmp.ReplyTo
//}
//
//func (t *Task) UUID() uuid.UUID {
//	return t.Header.Id
//}
//
//func (t *Task) Kwargs() map[string]any {
//	return t.Body.Kwargs
//}
//
//func (t *Task) Args() []any {
//	return t.Body.Args
//}
//
//func (t *Task) Reject() {
//	err := t.tmp.Reject(false)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func (t *Task) Nack() {
//	err := t.tmp.Nack(false, true)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// Тут нужно возвращать результат перед подтверждением и подумать, что делать при ошибке
//func (t *Task) Ack() {
//	err := t.tmp.Ack(false)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func (t *Task) Name() string {
//	return t.Header.Task
//}
//
//func NewTask(rawTask amqp.Delivery) *Task {
//	body, err := amqp_protocol.ParseTask(rawTask.Body)
//	if err != nil {
//		panic(err)
//	}
//	return &Task{
//		tmp:    rawTask,
//		Header: amqp_protocol.ParseHeader(rawTask.Headers),
//		Body:   body,
//	}
//}

func NewTask(rawTask amqp.Delivery) *baseTask.CeleryTask {
	body, err := amqp_protocol.ParseTask(rawTask.Body)
	if err != nil {
		panic(err)
	}
	return &baseTask.CeleryTask{
		TaskName:      "",
		TaskID:        uuid.UUID{},
		CorrelationID: uuid.UUID{},
		Args:          nil,
		Kwargs:        nil,
		Retries:       0,
		ETA:           0,
		Expires:       0,
		Callbacks:     "",
		Errbacks:      "",
	}
}
