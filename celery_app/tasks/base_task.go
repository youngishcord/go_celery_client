package base_tasks

import (
	amqp_protocol "celery_client/celery_app/core/implementations/rabbitmq/protocol"
	interf "celery_client/celery_app/core/interfaces"
	"fmt"

	"github.com/google/uuid"
)

type BaseTasks interface {
	Run() (any, error)
	Message() (any, error)
	Complete(any) // Метод завершения задачи
	UUID() uuid.UUID
	ReplyTo() string
}

type TaskConstructor func(message map[string]interface{}) (BaseTasks, error)

type BaseTask struct {
	name   string              `json:"name,omitempty"`
	args   []any               `json:"args,omitempty"`
	kwargs map[string]any      `json:"kwargs,omitempty"`
	embed  amqp_protocol.Embed `json:"embed,omitempty"`

	rawTask interf.Tasks
}

func (t *BaseTask) Complete(result any) {
	fmt.Println("Task complete call")
	fmt.Println(result)

	t.rawTask.Ack()
}

func (t *BaseTask) UUID() uuid.UUID {
	return t.rawTask.UUID()
}

func (t *BaseTask) ReplyTo() string {
	return t.rawTask.ReplyTo()
}

func NewBaseTask(rawTask interf.Tasks) BaseTask {
	return BaseTask{
		name:    rawTask.Name(),
		rawTask: rawTask,
	}
}
