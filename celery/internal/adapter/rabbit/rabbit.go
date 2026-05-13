package adapter

import (
	"go_celery_client/celery/config"
	rabbit "go_celery_client/celery/pkg/broker/rabbit"
	"go_celery_client/celery/protocol"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitAdapter struct {
	Client *rabbit.Client

	ConsumeCh *amqp.Channel
	PublishCh *amqp.Channel

	taskCh   chan *protocol.CeleryTask
	resultCh chan any // TODO: Celery like result need
}

func NewRabbitAdapter(client *rabbit.Client, settings config.BrokerSettings) (*RabbitAdapter, error) {
	adapter := RabbitAdapter{
		Client:   client,
		taskCh:   make(chan *protocol.CeleryTask),
		resultCh: make(chan any),
	}

	sub, err := client.Conn().Channel()
	if err != nil {
		log.Fatalln("NO_RABBITMQ_CONSUMER_CHANNEL_OPEN")
	}
	adapter.ConsumeCh = sub

	pub, err := client.Conn().Channel()
	if err != nil {
		log.Fatalln("NO_RABBITMQ_PUBLISHER_CHANNEL_OPEN")
	}
	adapter.PublishCh = pub

	err = sub.Qos(
		settings.Qos.PrefetchCount, // prefetch count
		settings.Qos.PrefetchSize,  // prefetch size (0 means unlimited)
		settings.Qos.Global,        // global (false = per consumer, true = per channel)
	)
	if err != nil {
		log.Fatalln("BAD_QOS_SETTINGS")
	}

	return &adapter, nil
}
