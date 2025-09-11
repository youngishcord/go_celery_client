package interfaces

type Tasks interface {
	Name() string

	Args() []any
	Kwargs() map[string]any

	//Payload() []byte

	Ack()
	Nack()
	Reject()
}
