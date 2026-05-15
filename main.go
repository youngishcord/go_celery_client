package main

import (
	"context"
	"go_celery_client/celery/pkg/logger"
	"log/slog"
	"sync"
)

func main() {
	logger := logger.NewLogger()
	ctx := context.Background()

	logger.InfoContext(ctx, "service started", slog.String("env", "dev"))

	logger.Info("text", slog.Any("data", map[string]any{
		"env":  "prod",
		"test": 1234,
	}))

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				logger.InfoContext(ctx,
					"worker event",
					slog.Int("worker_id", id),
					slog.Int("step", j),
				)
				//time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	logger.Info("service stopped")

	//app, err := celery_app.NewCeleryApp(config.CeleryConfig{
	//	Broker: config.BrokerSettings{},
	//	Worker: config.WorkerSettings{
	//		WorkerConcurrency: 2,
	//	},
	//	Queues: nil,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = app.RegisterTask("add", example.NewAddTask)
	//if err != nil {
	//	log.Fatalln("Failed task registration: ", err)
	//}
	//
	//// TODO: graceful shutdown and stop chan
	//err = app.Start()
	//if err != nil {
	//	return
	//}
	//
	//time.Sleep(2 * time.Second)
	//execCh := make(chan struct{})
	//<-execCh
}
