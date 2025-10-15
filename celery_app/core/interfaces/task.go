package interfaces

// type Task interface {
// 	Name() string

// 	Args() []any
// 	Kwargs() map[string]any

// 	UUID() uuid.UUID
// 	ReplyTo() string
// 	CorrelationID() string
// 	//Payload() []byte

// 	Ack()
// 	Nack()
// 	Reject()
// }

type Task interface {
	Run() (any, error)
	Message() (any, error)
	// Complete(any) // Метод завершения задачи
	// UUID() uuid.UUID
	// ReplyTo() string
	// CorrelationID() string
	// TaskChain() // TODO: тут нужен метод получения цепочки
}
