package broker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	conn        *amqp.Connection
	notifyChan  chan *amqp.Error
	notifyBlock chan amqp.Blocking

	config Config
}

func NewClient(config Config) (*Client, error) {
	client := &Client{
		config: config,
	}

	conn, err := amqp.Dial(config.Url())
	if err != nil {
		log.Fatalln("NO_RABBITMQ_CONNECTION")
	}
	client.conn = conn

	client.notifyChan = conn.NotifyClose(make(chan *amqp.Error))
	client.notifyBlock = conn.NotifyBlocked(make(chan amqp.Blocking))

	return client, nil
}

func (c *Client) Conn() *amqp.Connection {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}
