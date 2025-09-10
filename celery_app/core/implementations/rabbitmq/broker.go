package rabbit

import "celery_client/celery_app/dto"

func (b *RabbitMQBroker) ConsumeTask() <-chan dto.CeleryRawTask {
	return b.RawTaskCh
}
