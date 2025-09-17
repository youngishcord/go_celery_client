package result

import (
	"github.com/google/uuid"
)

//type Serializable interface {
//	json.Marshaler
//	json.Unmarshaler
//}

// TODO: Возможно необходимо перенести это в dto
type CeleryResult struct {
	Status Status    `json:"status"`
	Result any       `json:"result"`
	Trace  string    `json:"traceback"`
	TaskID uuid.UUID `json:"task_id"`
}

func NewCeleryResult(status Status, result any, trace string, taskID uuid.UUID) CeleryResult {
	return CeleryResult{
		Status: status,
		Result: result,
		Trace:  trace,
		TaskID: taskID,
	}
}
