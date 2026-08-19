package app

import (
	"context"
	"go_celery_client/celery/config"
	"go_celery_client/celery/internal/adapter/rabbit"
	"go_celery_client/celery/internal/exceptions"
	rabbit "go_celery_client/celery/pkg/broker/rabbit"
	"go_celery_client/celery/protocol"
	"go_celery_client/celery/task"
	"sync"
)

// Broker отвечает за отправку и получение сообщений
type Broker interface {
	Consume() <-chan *protocol.CeleryTask
	Publish() error
	Start(queues []string) error
	Ack(*protocol.CeleryTask) error
}

// Backend отвечает за работу с результатами задач
type Backend interface {
	PublishResult(ctx context.Context, result any, celeryTask *protocol.CeleryTask) error
	PublishException(ctx context.Context, result any, celeryTask *protocol.CeleryTask, trace string) error
}

type WorkerPool struct {
	wg      sync.WaitGroup
	closeCh chan struct{}
}

type CeleryApp struct {
	broker  Broker
	backend Backend

	// TaskRegistry хранит список зарегистрированных задач для исполнения. Регистрируется по имени
	taskRegistry map[string]func(task *protocol.CeleryTask) (task.Task, error)
	workerPool   *WorkerPool

	exceptionRegistry map[string]exceptions.BaseException

	conf config.CeleryConfig

	mu sync.RWMutex
}

func NewCeleryApp(conf config.CeleryConfig) (CeleryApp, error) {
	rabbitClient, err := rabbit.NewClient(rabbit.Config{
		Host: "localhost",
		Port: "5672",
		User: rabbit.User{
			Username: "admin",
			Password: "admin",
		},
	})
	if err != nil {
		return CeleryApp{}, err
	}

	rabbitAdapter, err := adapter.NewRabbitAdapter(rabbitClient, config.BrokerSettings{
		Qos: config.Qos{
			PrefetchCount: conf.Worker.WorkerConcurrency,
		},
	})
	if err != nil {
		return CeleryApp{}, err
	}

	return CeleryApp{
		broker:       rabbitAdapter,
		backend:      rabbitAdapter, // FIXME: нужно распределение
		taskRegistry: map[string]func(task *protocol.CeleryTask) (task.Task, error){},
		conf:         conf,
		workerPool: &WorkerPool{
			closeCh: make(chan struct{}),
		},
	}, nil
}
