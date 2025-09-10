package interfaces

type Broker interface {
	// Connect(queues []string) error
	// TaskChannel() chan amqp.Delivery // Я только что понял, что этот интерфейс не
	// будет работать с Redis, поскольку он не универсален
	// Consume() (<-chan UniversalMessageCustomType)

	// Connection() amqp.Connection

	// TODO: стоит переименовать данный метод в нечто более подходящее
	// По итогу реализация отвечает за получение задач из очереди и складывание их в канал,
	// я просто возвращаю канал, для дальнейшего прослушиывания.
	// TODO: можно реализовать модель базового брокера, которая будет автоматически включать нужные каналы.
	ConsumeTask() <-chan Tasks // Функция получения сообщения от брокера
}
