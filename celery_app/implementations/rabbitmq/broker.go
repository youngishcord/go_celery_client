package rabbit

import "celery_client/celery_app/core/dto/protocol"

func (b *RabbitMQ) ConsumeTask() <-chan protocol.CeleryTask {
	return b.TaskCh
}

// Ack basic acknowledge function for celery task
func (b *RabbitMQ) Ack(task task.CeleryTask) error {
	err := b.Consumer.Ack(task.DeliveryTag, false)
	if err != nil {
		return err
	}
	return nil
}

// Reject basic reject function for celery task
func (b *RabbitMQ) Reject(task task.CeleryTask, requeue bool) error {
	err := b.Consumer.Reject(task.DeliveryTag, requeue)
	if err != nil {
		return err
	}
	return nil
}

func (b *RabbitMQ) Nack(task task.CeleryTask, requeue bool) error {
	err := b.Consumer.Nack(task.DeliveryTag, false, requeue)
	if err != nil {
		return err
	}
	return nil
}
