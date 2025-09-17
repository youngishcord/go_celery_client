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
	Status Status
	Result any
	Trace  string
	TaskID uuid.UUID
}

func NewCeleryResult(status Status, result any, trace string, taskID uuid.UUID) CeleryResult {
	return CeleryResult{
		Status: status,
		Result: result,
		Trace:  trace,
		TaskID: taskID,
	}
}
