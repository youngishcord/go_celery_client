package main

import (
	"go_celery_client/celery/app"
	"go_celery_client/celery/config"
	example "go_celery_client/celery/examples/tasks"
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

	app.RegisterTask("add", example.NewAddTask)

	// TODO: graceful shutdown and stop chan
	err = app.Start()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)
	execCh := make(chan struct{})
	<-execCh
}
