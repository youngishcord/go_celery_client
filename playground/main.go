package main

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5545/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	//pub, err := conn.Channel()
	//if err != nil {
	//	panic(err)
	//}

	msgs, err := ch.Consume(
		"qwer",
		// TODO: тут надо сделать кастомное имя для консюмера из конфигурации
		fmt.Sprintf("consumer_"), // index
		false,                    // TODO: autoAck должен быть false по идее
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	msg := <-msgs
	fmt.Println(string(msg.Body))

	time.Sleep(20 * time.Second)

	err = ch.Ack(msg.DeliveryTag, false) // Ура так работает!
	if err != nil {
		panic(err)
	}

	//err = pub.Ack(msg.DeliveryTag, false) // Если канал не тот, из которого
	//if err != nil {                       // пришло сообщение, то не работает!
	//	panic(err)
	//}
}
