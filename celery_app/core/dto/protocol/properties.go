package protocol

import "github.com/google/uuid"

type DeliveryInfo struct {
	Exchange   string `json:"exchange,omitempty"`
	RoutingKey string `json:"routing_key,omitempty"`
}

type Properties struct {
	CorrelationID uuid.UUID    `json:"correlation_id"`
	ReplyTo       uuid.UUID    `json:"reply_to"`
	DeliveryMode  uint8        `json:"delivery_mode"`
	Priority      uint8        `json:"priority"`
	DeliveryInfo  DeliveryInfo `json:"delivery_info"`

	DeliveryTag  *uint64 `json:"delivery_tag,omitempty"`
	BodyEncoding *string `json:"body_encoding,omitempty"` // content encoding
}
