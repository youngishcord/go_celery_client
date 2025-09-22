package amqp_protocol

import "github.com/google/uuid"
import s "celery_client/celery_app/core/dto"

type AMQPCeleryResults struct {
	Status s.Status  `json:"status"`
	Result any       `json:"result"`
	Trace  string    `json:"traceback"`
	TaskID uuid.UUID `json:"task_id"`
}

func NewCeleryResult(status s.Status, result any, trace string, taskID uuid.UUID) AMQPCeleryResults {
	return AMQPCeleryResults{
		Status: status,
		Result: result,
		Trace:  trace,
		TaskID: taskID,
	}
}
