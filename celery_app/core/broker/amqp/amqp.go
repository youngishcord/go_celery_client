package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPBroker struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	Host string
	Port string

	user string
	pass string

	RawTaskCh chan amqp.Delivery
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

func (b *AMQPBroker) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
}

func (b *AMQPBroker) Connect(queues []string) error {
	conn, err := amqp.Dial(b.url())
	if err != nil {
		panic("NO RABBITMQ CONNECTION")
	}
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

	// Это надо как то вынести в отдельное место
	err = ch.Qos(
		2,     // prefetch count
		0,     // prefetch size (0 means unlimited)
		false, // global (false = per consumer, true = per channel)
	)
	if err != nil {
		panic("BAD QOS SETTINGS")
	}

	for index, queue := range queues {
		go func(queue string, ch *amqp.Channel) {
			// TODO: Нужен контекст?
			// TODO: Надо сделать проверку что очередь существует
			msgs, err := ch.Consume(queue,
				fmt.Sprintf("consumer_%d", index), // index
				true,
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
