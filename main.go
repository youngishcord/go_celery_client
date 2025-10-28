package main

import (
	celery "celery_client/celery_app/app"
	conf "celery_client/celery_app/celery_conf"
	tasks "celery_client/celery_app/tasks"
	"fmt"
)

func main() {

	app := celery.NewCeleryApp(conf.CeleryConf{
		Broker: conf.BrokerSettings{
			BrokerType: "RabbitMQ",
			ConnectionData: conf.Connection{
				Host: "localhost",
				Port: "5545",
				User: "guest",
				Pass: "guest",
			},
		},
		Backend: conf.BackendSettings{
			BackendType:    "RPC", //"Redis", //
			ConnectionData: conf.Connection{},
		},
		Worker: conf.WorkerSettings{
			WorkerConcurrency: 2,
		},
		Queues: []string{"qwer", "asdf"},
	})

	err := app.RegisterTask("add", tasks.NewAddTask) // тут передается конструктор, который дергается каждый раз при получении задачи.
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
}
