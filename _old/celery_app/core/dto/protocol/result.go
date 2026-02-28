package protocol

import (
	dto "celery_client/celery_app/core/dto"

	"github.com/google/uuid"
)

type CeleryResult struct {
	Status dto.Status `json:"status"`
	Result any        `json:"result"`
	Trace  string     `json:"trace"`
	TaskID uuid.UUID  `json:"task_id"`
}

func NewCeleryResult(status dto.Status, result any, trace string, taskID uuid.UUID) CeleryResult {
	return CeleryResult{
		Status: status,
		Result: result,
		Trace:  trace,
		TaskID: taskID,
	}
}
