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
	err = ch.Qos(
		// prefetch count Этот параметр не даст мне получать больше задач, чем сейчас выполняются,
		// следовательно, данный параметр должен быть настроен по количеству обработчиков
		2,
		0,     // prefetch size (0 means unlimited)
		false, // global (false = per consumer, true = per channel)
	)

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
	//msg := <-msgs
	//fmt.Println(string(msg.Body))

	tasks := make(chan amqp.Delivery)

	go func() {
		for task := range tasks {
			println("worker1")
			println(string(task.Body))
			println("Sleeping...")
			time.Sleep(10 * time.Second)
			println("next")
			task.Ack(false)
		}
	}()

	go func() {
		for task := range tasks {
			println("worker2")
			println(string(task.Body))
			println("Sleeping...")
			time.Sleep(8 * time.Second)
			println("next")
			task.Ack(false)
		}
	}()

	for msg := range msgs {
		tasks <- msg
	}

	//time.Sleep(10 * time.Second)

	//err = ch.Ack(msg.DeliveryTag, false) // Ура так работает!
	//if err != nil {
	//	panic(err)
	//}

	//err = pub.Ack(msg.DeliveryTag, false) // Если канал не тот, из которого
	//if err != nil {                       // пришло сообщение, то не работает!
	//	panic(err)
	//}

	//err = ch.Reject(msg.DeliveryTag, true)
	//if err != nil {
	//	panic(err)
	//}

	//err = ch.Nack(msg.DeliveryTag, false, false)
	//if err != nil {
	//	panic(err)
	//}
}
