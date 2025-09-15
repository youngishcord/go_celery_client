package interfaces

import (
	celery "celery_client/celery_app/core/message/result"
	//tasks "celery_client/celery_app/tasks"
)

type Backend interface {
	// FIXME: тут возникает циклический импорт при попытке передать интерфейс задачи,
	//  поскольку интерфейс задачи включает базовый интерфейс задачи, который лежит в пакете с интерфейсами.
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
