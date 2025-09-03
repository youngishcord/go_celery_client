package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	dto "celery_client/dto/header"
	app "celery_client/worker"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// CeleryTask Структура задачи Celery (упрощённая)
type CeleryTask struct {
	ID      string        `json:"id"`
	Task    string        `json:"task"`
	Args    []interface{} `json:"args"`
	Kwargs  interface{}   `json:"kwargs"`
	ReplyTo string        `json:"reply_to"`
}

// CeleryResult Структура результата (для backend rpc)
type CeleryResult struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
	Trace  string      `json:"traceback"`
	TaskID string      `json:"task_id"`
}

func main() {
	app.TestF()
	return

	fmt.Println("Begin")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5545/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	err = ch.Qos(
		2,     // prefetch count
		0,     // prefetch size (0 means unlimited)
		false, // global (false = per consumer, true = per channel)
	)
	if err != nil {
		log.Fatal(err)
	}
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
		// Указание очереди
		"qwer",       // queue
		"consumer_1", // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")
	c := 0
	for message := range msgs {
		// tmp, ok := message.Headers.(map[string]interface{})
		header, err := dto.MakeHeaderFromMap(message.Headers)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		if header.Task != "test" {
			err = message.Nack(false, true)
			if err != nil {
				log.Fatalln(err)
			} else {
				fmt.Println("task nack")
			}
			c++
			fmt.Println(c)
			continue
		}

		fmt.Println(header)
		fmt.Println("Received a message:", message)
		fmt.Println("Received a message body:", string(message.Body))
		//fmt.Println("Received a message task:", string(message.))

		var raw []interface{}
		if err := json.Unmarshal(message.Body, &raw); err != nil {
			log.Fatal(err)
			//err := message.Nack(false, true)
			//if err != nil {
			//	return
			//}
			//continue
		}
		fmt.Println(raw)
		fmt.Println(raw[0])

		//var tr CeleryTask
		//err := json.Unmarshal(message.Body, &tr)
		//if err != nil {
		//	log.Printf("Error unmarshalling task result: %s", err)
		//}
		//fmt.Println(tr)
		//
		res := CeleryResult{
			Status: "SUCCESS",
			Result: 123,
			Trace:  "",
			TaskID: message.CorrelationId,
		}

		body, err := json.Marshal(res)
		if err != nil {
			log.Printf("Error marshalling task result: %s", err)
		}

		// Слип для проверки работы временный ack очередей
		fmt.Println("sleep...")
		time.Sleep(2 * time.Second)
		fmt.Println("wakeup!")

		if message.ReplyTo != "" {
			err = ch.PublishWithContext(
				context.Background(),
				"",
				message.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: message.CorrelationId,
					Body:          body,
				},
			)
		}
		if err != nil {
			log.Printf("Error publishing task result: %s", err)
		}
		err = message.Ack(false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("task done")
	}
}
