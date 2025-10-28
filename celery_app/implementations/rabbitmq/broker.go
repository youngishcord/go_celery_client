package rabbit

import protocol "celery_client/celery_app/core/dto/protocol"

func (b *Rabbit) ConsumeTask() <-chan protocol.CeleryTask {
	return b.TaskCh
}

// Ack basic acknowledge function for celery task
func (b *Rabbit) Ack(task protocol.CeleryTask) error {
	err := b.Consumer.Ack(task.Properties.DeliveryTag, false)
	if err != nil {
		return err
	}
	return nil
}

// Reject basic reject function for celery task
func (b *Rabbit) Reject(task protocol.CeleryTask, requeue bool) error {
	err := b.Consumer.Reject(task.Properties.DeliveryTag, requeue)
	if err != nil {
		return err
	}
	return nil
}

// Nack negatively acknowledges a delivery by its delivery tag.
func (b *Rabbit) Nack(task protocol.CeleryTask, requeue bool) error {
	err := b.Consumer.Nack(task.Properties.DeliveryTag, false, requeue)
	if err != nil {
		return err
	}
	return nil
}
