package main

import (
	app "go_celery_client/celery/app"
	"go_celery_client/celery/config"
	examples "go_celery_client/celery/examples/tasks"
	"go_celery_client/celery/pkg/logger"
	"log"
	"time"
)

func main() {
	_ = logger.NewLogger(
		logger.WithSetDefault(),
	)

	app, err := app.NewCeleryApp(config.CeleryConfig{
		Broker: config.BrokerSettings{},
		Worker: config.WorkerSettings{
			WorkerConcurrency: 2,
		},
		Queues: nil,
	})
	if err != nil {
		panic(err)
	}

	err = app.RegisterTask("add", examples.NewAddTask)
	if err != nil {
		log.Fatalln("Failed task registration: ", err)
	}

	err = app.RegisterTask("panic", examples.NewPanicTask)
	if err != nil {
		log.Fatalln("Failed task registration: ", err)
	}

	err = app.RegisterTask("inf", examples.NewInfTask)
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
