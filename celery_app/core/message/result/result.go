package result

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Serializable interface {
	json.Marshaler
	json.Unmarshaler
}

type CeleryResult struct {
	Status Status
	Result Serializable
	Trace  string
	TaskID uuid.UUID
}

func NewCeleryResult(status Status, result Serializable, trace string) {

}
