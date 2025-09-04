package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPBroker struct {
	Conn *amqp.Connection

	Host string
	Port string

	user string
	pass string

	RawTaskCh chan amqp.Delivery
}

func (b *AMQPBroker) url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.user, b.pass, b.Host, b.Port)
}

func (b *AMQPBroker) Connect(queues []string) {
	conn, err := amqp.Dial(b.url())
	if err != nil {
		panic("NO RABBITMQ CONNECTION")
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Fatal("BAD RABBITMQ CONNECTION CLOSE")
		}
	}(conn)

	b.Conn = conn

	for queue := range queues {
		ch, err := conn.Channel()

	}

}

func (b *AMQPBroker) TMP1() {

}
