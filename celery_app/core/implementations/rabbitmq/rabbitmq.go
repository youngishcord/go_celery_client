package rabbit

import (
	"fmt"
	"log"

	conf "celery_client/celery_app/celery_conf"
	q "celery_client/celery_app/core/implementations/rabbitmq/queue"
	"celery_client/celery_app/dto"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Структура, хранящее подключение к RabbitMQ
type RabbitMQBroker struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel

	Host string
	Port string

	RawTaskCh chan dto.CeleryRawTask
	ResultCh  chan string // Служит для возврата результатов, если используется RPC backend

	user string
	pass string
}

func (b *RabbitMQBroker) connect(conf conf.CeleryConf) error {
	conn, err := amqp.Dial(b.url())
	if err != nil {
		panic("NO RABBITMQ CONNECTION")
	}

	// TODO: надо придумать, где будут деферы для закрытия подключений
	//defer func(conn *amqp.Connection) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatal("BAD RABBITMQ CONNECTION CLOSE")
	//	}
	//}(conn)

	b.Conn = conn
	ch, err := conn.Channel()
	//defer func(ch *amqp.Channel) {
	//	err := ch.Close()
	//	if err != nil {
	//		log.Fatal("BAD RABBITMQ CHANNEL CLOSE")
	//	}
	//}(ch)
	if err != nil {
		panic("NO RABBITMQ CHANNEL OPEN")
	}
	b.Channel = ch

	// TODO: Это надо как то вынести в отдельное место
	// Это конфиг, который должен быть настраиваемый снаружи
	err = ch.Qos(
		2,     // prefetch count
		0,     // prefetch size (0 means unlimited)
		false, // global (false = per consumer, true = per channel)
	)
	if err != nil {
		panic("BAD QOS SETTINGS")
	}

	// Это только консьюмеры на каждую очередь
	for index, queue := range conf.Queues {
		b.declareQueue(*q.NewDefaultQueue(queue))

		go func(queue string, ch *amqp.Channel) {
			// TODO: Нужен контекст?
			// TODO: Надо сделать проверку что очередь существует
			msgs, err := ch.Consume(
				queue,
				// TODO: тут надо сделать кастомное имя для консюмера из конфигурации
				fmt.Sprintf("consumer_%d", index), // index
				true,                              // TODO: autoAck должен быть false по идее
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Fatal(err)
			}
			for d := range msgs {
				fmt.Println(string(d.Body))
				// b.RawTaskCh <- d
				b.RawTaskCh <- b.newCeleryTask(d)
			}
		}(queue, ch)
	}

	return nil
}

func NewAMQPBroker(conf conf.CeleryConf) *RabbitMQBroker {
	broker := &RabbitMQBroker{
		Host:      conf.Broker.ConnectionData.Host,
		Port:      conf.Broker.ConnectionData.Port,
		user:      conf.Broker.ConnectionData.User,
		pass:      conf.Broker.ConnectionData.Pass,
		RawTaskCh: make(chan dto.CeleryRawTask),
	}

	broker.connect(conf)

	return broker
}
