package queue

import amqp "github.com/rabbitmq/amqp091-go"

// Подробнее в документации https://www.rabbitmq.com/docs/queues
// На самом деле внутри уже есть структура для очереди, а это
// получается не очень нужная обертка. В дальнейшем, наверное можно
// попробовать переиспользовать уже готовое решение.
type Queue struct {
	Name       string     // reference name
	Durable    bool       // the queue will survive a broker restart
	AutoDelete bool       // queue that has had at least one consumer is deleted when last consumer unsubscribes
	Exclusive  bool       // used by only one connection and the queue will be deleted when that connection closes
	NoWait     bool       //
	Args       amqp.Table // не знаю что может лежать в аргсах
}

func NewDefaultQueue(name string) *Queue {
	return &Queue{
		Name:       name,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
}

func NewResultQueue() *Queue {
	return &Queue{
		Name:       "",
		Durable:    false,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args: amqp.Table{
			"x-expires": int32(86400000),
		},
	}
}
