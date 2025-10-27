package main

import (
	celery "celery_client/celery_app/app"
	"celery_client/celery_app/celery_conf"
	"celery_client/celery_app/core/dto"
	base_tasks "celery_client/celery_app/tasks"
	"fmt"
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

	app := celery.NewCeleryApp(celery_conf.CeleryConf{
		Broker: dto.BrokerDto{
			BrokerType: "RabbitMQ",
			ConnectionData: dto.Connection{
				Host: "localhost",
				Port: "5545",
				User: "guest",
				Pass: "guest",
			},
		},
		Backend: dto.BackendDto{
			BackendType:    "RPC", //"Redis", //
			ConnectionData: dto.Connection{},
		},
		Queues: []string{"qwer", "asdf"},
	})

	err := app.RegisterTask("add", base_tasks.NewAddTask) // тут передается конструктор, который дергается каждый раз при получении задачи.
	if err != nil {
		return
	}

	// app.StartMessageDriver()
	err = app.RunWorker()
	if err != nil {
		fmt.Println(err)
	}

	exitCh := make(chan int)
	<-exitCh

	//for {
	//	log.Println(<-app.TaskPoolCh)
	//}

	// rabbit := NewAMQPBroker("localhost", "5545", "guest", "guest")

	// err := rabbit.Connect([]string{"qwer", "asdf"})
	// if err != nil {
	// 	return
	// }

	// for {
	// 	log.Println(string((<-rabbit.RawTaskCh).Body))
	// }
}
