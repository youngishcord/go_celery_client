package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	conn   *amqp.Connection
	config Config
}

func NewClient(config Config) (*Client, error) {
	client := &Client{
		config: config,
	}

	conn, err := amqp.Dial(config.Url())
	if err != nil {
		panic("NO_RABBITMQ_CONNECTION")
	}
	client.conn = conn

	return client, nil
}

func (c *Client) Conn() *amqp.Connection {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}
