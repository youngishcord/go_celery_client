package main

import (
	app "go_celery_client/celery/app"
	"go_celery_client/celery/config"
	examples "go_celery_client/celery/examples/tasks"
	"go_celery_client/celery/pkg/logger"
	"go_celery_client/celery/protocol"
	"go_celery_client/celery/task"
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

	err = app.RegisterTasks(map[string]func(task *protocol.CeleryTask) (task.Task, error){
		"add":   examples.NewAddTask,
		"panic": examples.NewPanicTask,
		"inf":   examples.NewInfTask,
		"err":   examples.NewErrorTask,
	})

	// TODO: graceful shutdown and stop chan
	err = app.Start()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)
	execCh := make(chan struct{})
	<-execCh
}
