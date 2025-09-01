package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	fmt.Println("Begin")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5545/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	// q, err := ch.QueueDeclare(
	// 	"celery", // name
	// 	false,    // durable
	// 	false,    // delete when unused
	// 	false,    // exclusive
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )
	// failOnError(err, "Fail on declare a queue")

	msgs, err := ch.Consume(
		"celery", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {

		fmt.Println(
		// json.Unmarshal(d.Body),
		)
	}

}
