package task

import "github.com/google/uuid"

type CeleryTask struct {
	TaskName      string         `json:"taskName,omitempty"`
	TaskID        uuid.UUID      `json:"taskId"`
	CorrelationID uuid.UUID      `json:"correlationId"`
	Args          []any          `json:"args,omitempty"`
	Kwargs        map[string]any `json:"kwargs,omitempty"`
	Retries       int            `json:"retries,omitempty"`
	ETA           int            `json:"eta,omitempty"`
	Expires       int64          `json:"expires,omitempty"`
	Callbacks     string         `json:"callbacks,omitempty"`
	Errbacks      string         `json:"errbacks,omitempty"`

	DeliveryTag uint64 `json:"deliveryTag,omitempty"` // RabbitMQ tag for Ack

	//embed amqp_protocol.Embed `json:"embed,omitempty"` // Это надо убрать

}

func (t *CeleryTask) Ack(result any) {

}
