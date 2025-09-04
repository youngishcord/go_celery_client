package main

import (
	. "celery_client/celery_app/core/broker/amqp"
	"log"
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
	rabbit := NewAMQPBroker("localhost", "5545", "guest", "guest")

	err := rabbit.Connect([]string{"qwer", "asdf"})
	if err != nil {
		return
	}

	for {
		log.Println(string((<-rabbit.RawTaskCh).Body))
	}
}
