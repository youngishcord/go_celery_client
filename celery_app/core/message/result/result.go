package result

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Serializable interface {
	json.Marshaler
	json.Unmarshaler
}

// TODO: Возможно необходимо перенести это в dto
type CeleryResult struct {
	Status Status
	Result Serializable
	Trace  string
	TaskID uuid.UUID
}

func NewCeleryResult(status Status, result Serializable, trace string, taskID uuid.UUID) CeleryResult {
	return CeleryResult{
		Status: status,
		Result: result,
		Trace:  trace,
		TaskID: taskID,
	}
}
