package rabbit

import (
	"celery_client/celery_app/core/interfaces"
)

func (b *RabbitMQBroker) ConsumeTask() <-chan interfaces.Tasks {
	return b.RawTaskCh
}
