package celery_app

import (
	conf "celery_client/celery_app/celery_conf"
	"celery_client/celery_app/core/dto/protocol"
	e "celery_client/celery_app/core/errors"
	"celery_client/celery_app/core/exceptions"
	rabbit "celery_client/celery_app/implementations/rabbitmq"
	"celery_client/celery_app/implementations/redis_client"

	result "celery_client/celery_app/message/result"
	"fmt"
	"log"
)

type Broker interface {
	// Connect(queues []string) error
	// TaskChannel() chan amqp.Delivery // Я только что понял, что этот интерфейс не
	// будет работать с Redis, поскольку он не универсален
	// Consume() (<-chan UniversalMessageCustomType)

	// Connection() amqp.Connection

	// TODO: стоит переименовать данный метод в нечто более подходящее
	// По итогу реализация отвечает за получение задач из очереди и складывание их в канал,
	// я просто возвращаю канал, для дальнейшего прослушиывания.
	// TODO: можно реализовать модель базового брокера, которая будет автоматически включать нужные каналы.
	ConsumeTask() <-chan protocol.CeleryTask // Функция получения сообщения от брокера
}

type Backend interface {
	// FIXME: тут возникает циклический импорт при попытке передать интерфейс задачи,
	//  поскольку интерфейс задачи включает базовый интерфейс задачи, который лежит в пакете с интерфейсами.
	PublishResult(result any, celeryTask protocol.CeleryTask) error
	// FIXME: Мне не нравится,
	//  что тут разные интерфейсы у публикации результата и ошибки. Возможно стоит оставить
	//  только интерфейс сырой таски, поскольку он также имеет метод получения id задачи
	PublishException(result any, celeryTask protocol.CeleryTask, trace string) error
	ConsumeResult(taskID string) (<-chan result.CeleryResult, error)
}

type Task interface {
	Run() (any, error)
	Message() (any, error)
}

type CeleryApp struct {
	TasksRegistry map[string]func(task protocol.CeleryTask) (Task, error)
	// TaskPoolCh    chan BaseTasks

	ResultCh chan result.CeleryResult

	Broker  Broker  // Наверное структура или интерфейс, которая описывает подключение к брокеру
	Backend Backend // Наверное структура или интерфейс, которая описывает подключение к бекенду
	// думаю всетаки интерфейсы, поскольку и брокер и бекенд могут быть разными (redis и rabbit)

	appConf conf.CeleryConf
}

func (a *CeleryApp) RegisterTask(name string, constructor func(task protocol.CeleryTask) (Task, error)) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = constructor
	return nil
}

// GetTask Получение задачи из реестра.
// Может быть приватный?
func (a *CeleryApp) GetTask(name string) (Task, error) {
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

// RunWorker запуска воркеров в отдельных горутинах
func (a *CeleryApp) RunWorker() error {
	a.celeryStartupMessage()

	// TODO: Вынеси меня в переменную в структуре
	taskPool := a.Broker.ConsumeTask()

	//for i := 0; i < 1; i++ { // FIXME: Настроить по количеству воркеров в конфиге!
	go func() {
		for celeryTask := range taskPool {
			task, err := a.MakeTask(celeryTask)
			if err != nil {
				panic(err) // FIXME: Send error from worker
			}

			result, err := task.Run()
			if err != nil {
				panic(err) // FIXME: Send error from worker
			}

			err = a.Backend.PublishResult(result, celeryTask)
			if err != nil {
				panic(err) // FIXME: Send error from worker
			}
		}
	}()
	return nil
}

// Delay Отправка задачи в очередь
// Метод должен возвращать сущность задачи, которую можно поставить на
// ожидание и получить результат.
func (a *CeleryApp) Delay(task_name string, args []any, kwargs map[any]any) Task {
	panic("IMPLEMENT ME")
}

// Get Получение результата задачи по ее сущности из backend
func (a *CeleryApp) Get(task Task) result.CeleryResult {
	panic("IMPLEMENT ME")
}

// MakeTask получает на вход параметры, находит конструктор задачи, фармирует
// структуру и возвращает ее для дальнейшей обработки.
func (a *CeleryApp) MakeTask(task protocol.CeleryTask) (Task, error) {
	constructor, ok := a.TasksRegistry[task.Headers.Task]
	if !ok {
		_ = a.Backend.PublishException(
			exceptions.GetException(e.NotRegistered,
				[]string{task.Headers.Task}),
			task,
			"test trace",
		)
		log.Println("TASK NOT FOUND")
		return nil, e.NotRegistered
	}

	// Registered constructor function allready return error
	return constructor(task)

	// newTask, _ := f(task)

	// a.TaskPoolCh <- newTask

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

}

// FIXME: Переделать в приватный и вызывать в конструкторе
func (a *CeleryApp) StartMessageDriver() {
	rawTaskChannel := a.Broker.ConsumeTask()
	go func() {
		for rawTask := range rawTaskChannel {
			a.MakeTask(rawTask)
		}
	}()
}

func NewBrokerAndBackend(conf conf.CeleryConf) (Broker, Backend) {
	// TODO: расширение функциональности для работы с redis
	var broker Broker
	var backend Backend
	// var err error

	switch conf.Broker.BrokerType {
	case "RabbitMQ":
		broker = rabbit.NewAMQPBroker(conf)
	default:
		panic(fmt.Errorf("ЗНАЧЕНИЕ БРОКЕРА НЕ ОПРЕДЕЛЕНО"))
	}

	switch conf.Backend.BackendType {
	case "RPC":
		if tmp, ok := broker.(Backend); ok {
			backend = tmp
		}
	case "Redis":
		backend = redis_client.NewRedisClient() // TODO: тут нужны параметры
	default:
		panic(fmt.Errorf("ЗНАЧЕНИЕ БЕКЕНДА НЕ ОПРЕДЕЛЕНО"))
	}

	return broker, backend
}

func NewCeleryApp(conf conf.CeleryConf) *CeleryApp {

	broker, backend := NewBrokerAndBackend(conf)

	app := &CeleryApp{
		TasksRegistry: map[string]func(task protocol.CeleryTask) (Task, error){},
		ResultCh:      make(chan result.CeleryResult, 1),
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
