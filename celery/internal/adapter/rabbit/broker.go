package adapter

import (
	"fmt"
	"go_celery_client/celery/protocol"
	"log"
)
import queues "go_celery_client/celery/internal/adapter/rabbit/queue"

func (a *RabbitAdapter) Start(q []string) error {
	for index, queue := range q {
		tmp := queues.NewDefaultQueue(queue)
		_, err := a.ConsumeCh.QueueDeclare(
			tmp.Name,
			tmp.Durable,
			tmp.AutoDelete,
			tmp.Exclusive,
			tmp.NoWait,
			tmp.Args,
		)
		if err != nil {
			return err
		}

		go func() {
			msgs, err := a.ConsumeCh.Consume(
				queue,
				// TODO: тут надо сделать кастомное имя для консюмера из конфигурации
				fmt.Sprintf("consumer_%d", index), // index
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Fatalf("Failed to register a consumer: %s", err)
				return
			}

			for msg := range msgs {
				task, err := newTask(msg)
				if err != nil {
					msg.Nack(false, true)
					log.Fatalf("Failed to create a task: %s", err)
					return
				}

				if a.taskCh == nil {
					log.Fatalf("Task channel is nil")
				}
				a.taskCh <- task

				//select {
				//case msg
				//case <- time.After(5 * time.Second):
				//	msg.Nack(false, true)
				//	log.Println("Queue %s is backed up, nacking message", queue)
				//}
			}
		}()
	}
	return nil
}

func (a *RabbitAdapter) Ack(task *protocol.CeleryTask) error {
	return a.ConsumeCh.Ack(task.Properties.DeliveryTag, false)
}

func (a *RabbitAdapter) Consume() <-chan *protocol.CeleryTask {
	return a.taskCh
}

func (a *RabbitAdapter) Publish() error {
	return nil
}
