package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5545/")
	failOnError(err, "connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "open channel")
	defer ch.Close()

	// 1) объявляем exchange (например, direct)
	err = ch.ExchangeDeclare(
		"my_exchange", // имя exchange
		"direct",      // тип
		true,          // durable
		false,         // auto-delete
		false,         // internal
		false,         // no-wait
		nil,           // аргументы
	)
	failOnError(err, "declare exchange")

	// 2) объявляем очередь
	q, err := ch.QueueDeclare(
		"my_queue", // имя очереди
		true,       // durable — чтобы очередь пережила рестарт RabbitMQ
		false,      // auto-delete
		false,      // exclusive
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "declare queue")

	// 3) привязываем очередь к exchange с routing key
	err = ch.QueueBind(
		q.Name,           // имя очереди
		"my_routing_key", // routing key
		"my_exchange",    // exchange
		false,
		nil,
	)
	failOnError(err, "bind queue")

	// 4) Теперь очередь гарантированно существует — можно публиковать сообщения
	body := []byte(`{"foo": "bar"}`) // JSON, например — тело задачи Celery
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false, false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	failOnError(err, "publish")
	log.Println("Message published")

	log.Println(q.Name)

}
