package broker

import (
	"fmt"
	"log"

	q "celery_client/celery_app/core/broker/amqp/queue"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPBroker struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	Host string
	Port string

	RawTaskCh chan amqp.Delivery

	user string
	pass string
}

func NewAMQPBroker(host string, port string, user string, pass string) *AMQPBroker {
	return &AMQPBroker{
		Host:      host,
		Port:      port,
		user:      user,
		pass:      pass,
		RawTaskCh: make(chan amqp.Delivery),
	}
}

func (b *AMQPBroker) TaskChannel() chan amqp.Delivery {
	return b.RawTaskCh
}

func (b *AMQPBroker) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
}

func (b *AMQPBroker) declareQueue(queue q.Queue) {
	if b.Channel == nil {
		panic("CHANNEL NOT OPEN")
	}

	_, err := b.Channel.QueueDeclare(
		queue.Name,
		queue.Durable,
		queue.AutoDelete,
		queue.Exclusive,
		queue.NoWait,
		queue.Args,
	)
	if err != nil {
		panic("QUEUE WAS NOT DECLARED")
	}

}

func (b *AMQPBroker) Connect(queues []string) error {
	conn, err := amqp.Dial(b.url())
	if err != nil {
		panic("NO RABBITMQ CONNECTION")
	}

	// TODO: надо придумать, где будут деферы для закрыти подключений
	//defer func(conn *amqp.Connection) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatal("BAD RABBITMQ CONNECTION CLOSE")
	//	}
	//}(conn)

	b.Conn = conn
	ch, err := conn.Channel()
	//defer func(ch *amqp.Channel) {
	//	err := ch.Close()
	//	if err != nil {
	//		log.Fatal("BAD RABBITMQ CHANNEL CLOSE")
	//	}
	//}(ch)
	if err != nil {
		panic("NO RABBITMQ CHANNEL OPEN")
	}
	b.Channel = ch

	// TODO: Это надо как то вынести в отдельное место
	// Это конфиг, который должен быть настраиваемый снаружи
	err = ch.Qos(
		2,     // prefetch count
		0,     // prefetch size (0 means unlimited)
		false, // global (false = per consumer, true = per channel)
	)
	if err != nil {
		panic("BAD QOS SETTINGS")
	}

	// Это только консьюмеры на каждую очередь
	for index, queue := range queues {
		b.declareQueue(*q.NewDefaultQueue(queue))

		go func(queue string, ch *amqp.Channel) {
			// TODO: Нужен контекст?
			// TODO: Надо сделать проверку что очередь существует
			msgs, err := ch.Consume(
				queue,
				fmt.Sprintf("consumer_%d", index), // index
				true,                              // autoAck должен быть false по идее
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}
			for d := range msgs {
				//fmt.Println(string(d.Body))
				b.RawTaskCh <- d
			}
		}(queue, ch)
	}

	return nil
}
