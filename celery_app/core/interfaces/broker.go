package interfaces

import dto "celery_client/celery_app/dto"

type Broker interface {
	// Connect(queues []string) error
	// TaskChannel() chan amqp.Delivery // Я только что понял, что этот интерфейс не
	// будет работать с Redis, поскольку он не универсален
	// Consume() (<-chan UniversalMessageCustomType)

	// Connection() amqp.Connection

	ConsumeTask() <-chan dto.CeleryRawTask // Функция получения сообщения от брокера
}
