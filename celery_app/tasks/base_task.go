package base_tasks

import (
	"celery_client/celery_app/core/implementations/rabbitmq/protocol"
)

type BaseTasks interface {
	Run() (any, error)
	Message() (any, error)
	Complete() // Метод завершения задачи
}

type TaskConstructor func(message map[string]interface{}) (BaseTasks, error)

type BaseTask struct {
	name   string              `json:"name,omitempty"`
	args   []any               `json:"args,omitempty"`
	kwargs map[string]any      `json:"kwargs,omitempty"`
	embed  amqp_protocol.Embed `json:"embed,omitempty"`
}

func NewBaseTask(name string) BaseTask {
	return BaseTask{name: name}
}
