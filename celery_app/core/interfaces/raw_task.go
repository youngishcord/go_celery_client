package interfaces

import "github.com/google/uuid"

type Tasks interface {
	Name() string

	Args() []any
	Kwargs() map[string]any

	UUID() uuid.UUID
	ReplyTo() string
	//Payload() []byte

	Ack()
	Nack()
	Reject()
}
