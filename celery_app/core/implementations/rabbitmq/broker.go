package rabbit

import (
	"celery_client/celery_app/core/interfaces"
)

func (b *RabbitMQ) ConsumeTask() <-chan interfaces.Tasks {
	return b.RawTaskCh
}
