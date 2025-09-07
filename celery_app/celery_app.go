package celery_app

import (
	conf "celery_client/celery_app/celery_conf"
	b "celery_client/celery_app/core/broker"
	"log"

	amqpBroker "celery_client/celery_app/core/broker/amqp"
	. "celery_client/celery_app/core/message/amqp/protocol"
	. "celery_client/celery_app/core/message/result"
	. "celery_client/celery_app/tasks"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type CeleryApp struct {
	TasksRegistry map[string]func(message []byte) (BaseTasks, error)
	TaskPoolCh    chan BaseTasks
	ResultCh      chan any

	Broker      b.Broker // Наверное структура или интерфейс, которая описывает подключение к брокеру
	BackendConn string   // Наверное структура или интерфейс, которая описывает подключение к бекенду
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

// RunWorker Запуск основного треда воркера
func (a *CeleryApp) RunWorker() error {
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
	return nil
}

// Get Получение результата задачи по ее сущности из backend
func (a *CeleryApp) Get(task BaseTasks) CeleryResult {
	return CeleryResult{}
}

func (a *CeleryApp) MakeTask(task amqp.Delivery) {
	//task :=
	fmt.Println(task)
	header := MakeHeaderFromTable(task.Headers)
	//if err != nil {
	//	err := task.Nack(false, false)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	fmt.Println(header)

	fmt.Println("TASK NAME")
	fmt.Println(header.Task)

	//a.TaskPoolCh <-
	//task.Ack()

	f, ok := a.TasksRegistry[header.Task]
	if !ok {
		log.Println("TASK NOT FOUND")
		return
	}

	newTask, _ := f(task.Body)

	a.TaskPoolCh <- newTask

	return
}

func (a *CeleryApp) StartMessageDriver() {
	go func() {
		for rawTask := range a.Broker.TaskChannel() {
			a.MakeTask(rawTask)
		}
	}()
}

func NewCeleryApp(conf conf.CeleryConf) *CeleryApp {

	app := &CeleryApp{
		TasksRegistry: map[string]func(message []byte) (BaseTasks, error){},
		TaskPoolCh:    make(chan BaseTasks, 5), // по количеству запускаемых воркеров?
		ResultCh:      make(chan any),
		Broker: amqpBroker.NewAMQPBroker(
			conf.Broker.Host,
			conf.Broker.Port,
			conf.Broker.User,
			conf.Broker.Pass,
		),
		appConf:     conf,
		BackendConn: "",
	}

	err := app.Broker.Connect(conf.Queues)
	if err != nil {
		return nil
	}
	//app.TaskPoolCh <- app.Broker.TaskChannel()

	return app
}
