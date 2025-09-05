package celery_app

import (
	conf "celery_client/celery_app/celery_conf"
	b "celery_client/celery_app/core/broker"

	amqpBroker "celery_client/celery_app/core/broker/amqp"
	. "celery_client/celery_app/core/message/result"
	. "celery_client/celery_app/tasks"
	"fmt"
)

type CeleryApp struct {
	TasksRegistry map[string]BaseTasks
	TaskPoolCh    chan BaseTasks
	ResultCh      chan any

	Broker      b.Broker // Наверное структура или интерфейс, которая описывает подключение к брокеру
	BackendConn string   // Наверное структура или интерфейс, которая описывает подключение к бекенду
	// думаю всетаки интерфейсы, поскольку и брокер и бекенд могут быть разными (redis и rabbit)

	appConf conf.CeleryConf
}

func (a *CeleryApp) RegisterTask(name string, task BaseTasks) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = task
	return nil
}

// Получение задачи из реестра
// Может быть приватный?
func (a *CeleryApp) GetTask(name string) (BaseTasks, error) {
	return nil, nil
}

// Запуск основного треда воркера
func (a *CeleryApp) RunWorker() error {
	return nil
}

// Отправка задачи в очередь
// Метод должен возвращать сущность задачи, которую можно поставить на
// ожидание и получить результат.
func (a *CeleryApp) Delay(task_name string, args []any, kwargs map[any]any) BaseTasks {
	return nil
}

// Получение результата задачи по ее сущности из backend
func (a *CeleryApp) Get(task BaseTasks) CeleryResult {
	return CeleryResult{}
}

func NewCeleryApp(conf conf.CeleryConf) *CeleryApp {

	app := &CeleryApp{
		TasksRegistry: map[string]BaseTasks{},
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

	app.Broker.Connect(conf.Queues)
	app.TaskPoolCh <- app.Broker.TaskChannel()

	return app
}
