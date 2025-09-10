package interfaces

import (
	celery "celery_client/celery_app/core/message/result"
)

type Backend interface {
	// PublishResult(taskID string, body []byte, headers map[string]any) error
	PublishResult(result celery.CeleryResult) error
	ConsumeResult(taskID string) (<-chan celery.CeleryResult, error)
}

// конструктор отвечает за подключение,

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
