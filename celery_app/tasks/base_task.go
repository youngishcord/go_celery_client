package base_tasks

import (
	interf "celery_client/celery_app/core/interfaces"
	"celery_client/celery_app/implementations/rabbitmq/protocol"
	"fmt"

	"github.com/google/uuid"
)

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

func (t *BaseTask) CorrelationID() string {
	return t.rawTask.CorrelationID()
}

func (t *BaseTask) 


func NewBaseTask(rawTask interf.Tasks) BaseTask {
	return BaseTask{
		name:    rawTask.Name(),
		rawTask: rawTask,
	}
}

