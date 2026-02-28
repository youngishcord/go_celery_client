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

	// Delay отправляет задачу в очередь на исполнение и возвращает ее идентификатор
	// Delay() string
}

type Backend interface {
	// FIXME: тут возникает циклический импорт при попытке передать интерфейс задачи,
	//  поскольку интерфейс задачи включает базовый интерфейс задачи, который лежит в пакете с интерфейсами.
	PublishResult(result any, celeryTask protocol.CeleryTask) error

	// PublishException FIXME: Мне не нравится,
	//  что тут разные интерфейсы у публикации результата и ошибки. Возможно стоит оставить
	//  только интерфейс сырой таски, поскольку он также имеет метод получения id задачи
	PublishException(result any, celeryTask protocol.CeleryTask, trace string) error
	ConsumeResult(taskID string) (<-chan result.CeleryResult, error)

	// Get метод получения результатов отправленной в очередь задачи
	// Get()
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

	taskCh <-chan protocol.CeleryTask

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
	a.startMessageDriver()
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

	// Registered constructor function already return error
	return constructor(task)
}

// FIXME: Переделать в приватный и вызывать в конструкторе
func (a *CeleryApp) startMessageDriver() {
	a.taskCh = a.Broker.ConsumeTask()

	for i := 0; i < a.appConf.Worker.WorkerConcurrency; i++ {
		go func(counter int) {
			for celeryTask := range a.taskCh {
				task, err := a.MakeTask(celeryTask)
				if err != nil {
					panic(err) // FIXME: Send error from worker
				}

				taskResult, err := task.Run()
				if err != nil {
					panic(err) // FIXME: Send error from worker
				}

				err = a.Backend.PublishResult(taskResult, celeryTask)
				if err != nil {
					panic(err) // FIXME: Send error from worker
				}
			}
		}(i)
	}
}

func NewBrokerAndBackend(conf conf.CeleryConf) (Broker, Backend) {
	var broker Broker
	var backend Backend

	switch conf.Broker.BrokerType {
	case "RabbitMQ":
		r := rabbit.NewAMQPBroker(conf)
		broker = r
		backend = r
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
	return app
}
