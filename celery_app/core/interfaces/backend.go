package interfaces

import (
	celery "celery_client/celery_app/message/result"
)

type Backend interface {
	// FIXME: тут возникает циклический импорт при попытке передать интерфейс задачи,
	//  поскольку интерфейс задачи включает базовый интерфейс задачи, который лежит в пакете с интерфейсами.
	PublishResult(result any, baseTasks BaseTasks) error
	// FIXME: Мне не нравится,
	//  что тут разные интерфейсы у публикации результата и ошибки. Возможно стоит оставить
	//  только интерфейс сырой таски, поскольку он также имеет метод получения id задачи
	PublishException(result any, baseTasks Tasks, trace string) error
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
