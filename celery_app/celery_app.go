package celery_app

import (
	conf "celery_client/celery_app/celery_conf"
	rabbit "celery_client/celery_app/core/implementations/rabbitmq"
	interf "celery_client/celery_app/core/interfaces"
	. "celery_client/celery_app/core/message/result"
	. "celery_client/celery_app/tasks"
	"fmt"
)

type CeleryApp struct {
	TasksRegistry map[string]func(message []byte) (BaseTasks, error)
	TaskPoolCh    chan BaseTasks
	ResultCh      chan any

	Broker  interf.Broker  // Наверное структура или интерфейс, которая описывает подключение к брокеру
	Backend interf.Backend // Наверное структура или интерфейс, которая описывает подключение к бекенду
	// думаю всетаки интерфейсы, поскольку и брокер и бекенд могут быть разными (redis и rabbit)

	appConf conf.CeleryConf
}

func (a *CeleryApp) RegisterTask(name string, constructor func(message []byte) (BaseTasks, error)) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = constructor
	return nil
}

// GetTask Получение задачи из реестра.
// Может быть приватный?
func (a *CeleryApp) GetTask(name string) (BaseTasks, error) {
	return nil, nil
}

func (a *CeleryApp) celeryStartupMessage() {
	fmt.Println("Hello from CeleryApp")
	fmt.Printf("App listen %d queues: ", len(a.appConf.Queues))
	for _, q := range a.appConf.Queues {
		fmt.Printf("%s, ", q)
	}
	fmt.Println("\nRegistered tasks:")
	for name := range a.TasksRegistry {
		fmt.Println("\t.", name)
	}
}

// RunWorker Запуск основного треда воркера
func (a *CeleryApp) RunWorker() error {
	a.celeryStartupMessage()

	//for i := 0; i < 1; i++ {
	go func() {
		for task := range a.TaskPoolCh {
			result, err := task.Run()
			if err != nil {
				fmt.Println("RESULT ERROR")
				return
			}
			fmt.Println("THIS IS RESULT")
			fmt.Println(result)
		}
	}()
	//}
	return nil
}

// Delay Отправка задачи в очередь
// Метод должен возвращать сущность задачи, которую можно поставить на
// ожидание и получить результат.
func (a *CeleryApp) Delay(task_name string, args []any, kwargs map[any]any) BaseTasks {
	panic("IMPLEMENT ME")
}

// Get Получение результата задачи по ее сущности из backend
func (a *CeleryApp) Get(task BaseTasks) CeleryResult {
	panic("IMPLEMENT ME")
}

func (a *CeleryApp) MakeTask(task interf.Tasks) {
	//task :=
	fmt.Println(task)

	// header := MakeHeaderFromTable(task.Headers)

	//if err != nil {
	//	err := task.Nack(false, false)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	// fmt.Println(header)

	// fmt.Println("TASK NAME")
	// fmt.Println(header.Task)

	//a.TaskPoolCh <-
	//task.Ack()

	// f, ok := a.TasksRegistry[header.Task]
	// if !ok {
	// 	log.Println("TASK NOT FOUND")
	// 	return
	// }

	// newTask, _ := f(task.Body)

	// a.TaskPoolCh <- newTask

}

func (a *CeleryApp) StartMessageDriver() {
	rawTaskChannel := a.Broker.ConsumeTask()
	go func() {
		for rawTask := range rawTaskChannel {
			a.MakeTask(rawTask)
		}
	}()
}

func NewBrokerAndBackend(conf conf.CeleryConf) (interf.Broker, interf.Backend) {
	// TODO: расширение функциональности для работы с redis
	var broker interf.Broker
	var backend interf.Backend
	// var err error

	switch conf.Broker.BrokerType {
	case "RabbitMQ":
		broker = rabbit.NewAMQPBroker(conf)
	}

	switch conf.Backend.BackendType {
	case "RPC":
		if tmp, ok := broker.(interf.Backend); ok {
			backend = tmp
		}
	}

	return broker, backend
}

func NewCeleryApp(conf conf.CeleryConf) *CeleryApp {

	broker, backend := NewBrokerAndBackend(conf)

	app := &CeleryApp{
		TasksRegistry: map[string]func(message []byte) (BaseTasks, error){},
		TaskPoolCh:    make(chan BaseTasks, 5), // по количеству запускаемых воркеров?
		ResultCh:      make(chan any),
		Broker:        broker,
		Backend:       backend,
		appConf:       conf,
	}

	// err := app.Broker.Connect(conf.Queues)
	// if err != nil {
	// 	return nil
	// }

	// switch conf.Backend.BackendType {
	// case "RPC":
	// 	connection := app.Broker.Connection()
	// 	channel, err := connection.Channel()
	// 	if err != nil {
	// 		return nil
	// 	}
	// 	backend.NewAMQPBackend(channel)
	// }

	//app.TaskPoolCh <- app.Broker.TaskChannel()

	return app
}
