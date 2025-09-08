package backend

type Backend interface {
	MakeResult() error
	Connect() error
}

//func NewBackend(backend dto.BackendDto) Backend {
//	switch backend.BackendType {
//	case "RPC":
//		return &amqp.AMQPBackend{
//			RPCChannel: nil,
//		}
//	default:
//		panic("Unsupported backend type: " + backend.BackendType)
//	}
//}
