package base_tasks

import (
	amqp_protocol "celery_client/celery_app/core/implementations/rabbitmq/protocol"
	interf "celery_client/celery_app/core/interfaces"
	results "celery_client/celery_app/core/message/result"
	"fmt"
)

type BaseTasks interface {
	Run() (results.CeleryResult, error)
	Message() (any, error)
	Complete(results.CeleryResult) // Метод завершения задачи
}

type TaskConstructor func(message map[string]interface{}) (BaseTasks, error)

type BaseTask struct {
	name   string              `json:"name,omitempty"`
	args   []any               `json:"args,omitempty"`
	kwargs map[string]any      `json:"kwargs,omitempty"`
	embed  amqp_protocol.Embed `json:"embed,omitempty"`

	rawTask interf.Tasks
}

func (t *BaseTask) Complete(result results.CeleryResult) {
	fmt.Println("Task complete call")
	t.rawTask.Ack()
}

func NewBaseTask(rawTask interf.Tasks) BaseTask {
	return BaseTask{
		name:    rawTask.Name(),
		rawTask: rawTask,
	}
}
