package rabbit

import (
	interf "celery_client/celery_app/core/interfaces"
	r "celery_client/celery_app/message/result"
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Отношение к интерфейсу backend при работе с RPC
func (b *RabbitMQ) PublishResult(result r.CeleryResult, task interf.BaseTasks) error {
	// TODO: тут стоит задуматься над тем что будет, если во второй операции выпадет ошибка, а
	//  первая уже будет выполнена
	// TODO: Подтверждение результата должно быть другим. Я возвращаю тут результат
	//  сразу в нужное место используя amqp либу и только после успешной публикации делаю Ack
	//b.ResultCh <- result
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				err = ch.PublishWithContext(ctx,
					"",              // exchange
					resultQueue.Name, // routing key
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(result),
					})
				cancel()
				if err != nil {
					// если не удалось отправить результат — не подтверждаем задачу
					log.Printf("Ошибка при отправке результата: %v", err)
					continue
				}

				// подтверждаем, что задачу обработали
				if err := msg.Ack(false); err != nil {
					log.Printf("Ошибка при подтверждении: %v", err)
				} else {
					log.Printf("Задача '%s' подтверждена", task)
				}
	*/

	// TODO: Тут также нужна специальная обработка результата при работе с цепочкой
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := b.Channel.PublishWithContext(
		ctx,
		"",
		task.ReplyTo(),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("test result"),
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Отношение к интерфейсу backend при работе с RPC
func (b *RabbitMQ) ConsumeResult(taskID string) (<-chan r.CeleryResult, error) {
	// TODO: implement me
	panic("IMPLEMENT ME")
}
