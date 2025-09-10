package rabbit

import r "celery_client/celery_app/core/message/result"

// Отношение к интерфейсу backend при работе с RPC
func (b *RabbitMQBroker) PublishResult(result r.CeleryResult) error {
	// TODO: implement me
	panic("IMPLEMENT ME")
}

// Отношение к интерфейсу backend при работе с RPC
func (b *RabbitMQBroker) ConsumeResult(taskID string) (<-chan r.CeleryResult, error) {
	// TODO: implement me
	panic("IMPLEMENT ME")
}
