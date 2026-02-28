package protocol

import "github.com/google/uuid"

type DeliveryInfo struct {
	Exchange   string `json:"exchange,omitempty"`
	RoutingKey string `json:"routing_key,omitempty"`
}

type Properties struct {
	CorrelationID uuid.UUID `json:"correlation_id"`
	DeliveryTag   uint64    `json:"delivery_tag"`

	ReplyTo      uuid.UUID    `json:"reply_to,omitempty"`
	DeliveryMode uint8        `json:"delivery_mode,omitempty"`
	DeliveryInfo DeliveryInfo `json:"delivery_info,omitempty"`
	Priority     uint8        `json:"priority,omitempty"`
	BodyEncoding string       `json:"body_encoding,omitempty"` // content encoding
}
