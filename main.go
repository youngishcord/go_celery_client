package main

import (
	"go_celery_client/celery/app"
	"go_celery_client/celery/config"
	example "go_celery_client/celery/examples/tasks"
	"log"
	"time"
)

func main() {
	app, err := celery_app.NewCeleryApp(config.CeleryConfig{
		Broker: config.BrokerSettings{},
		Worker: config.WorkerSettings{
			WorkerConcurrency: 2,
		},
		Queues: nil,
	})
	if err != nil {
		panic(err)
	}

	err = app.RegisterTask("add", example.NewAddTask)
	if err != nil {
		log.Fatalln("Failed task registration: ", err)
	}

	// TODO: graceful shutdown and stop chan
	err = app.Start()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)
	execCh := make(chan struct{})
	<-execCh
}
