package rabbit

import (
	interf "celery_client/celery_app/core/interfaces"
	r "celery_client/celery_app/message/result"
	"context"
	"encoding/json"
	"fmt"
	"time"

	s "celery_client/celery_app/core/dto"
	protocol "celery_client/celery_app/implementations/rabbitmq/protocol"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Отношение к интерфейсу backend при работе с RPC
func (b *RabbitMQ) PublishResult(result any, task interf.BaseTasks) error {
	// TODO: тут стоит задуматься над тем что будет, если во второй операции выпадет ошибка, а
	//  первая уже будет выполнена
	// TODO: Подтверждение результата должно быть другим. Я возвращаю тут результат
	//  сразу в нужное место используя amqp либу и только после успешной публикации делаю Ack

	// TODO: Тут также нужна специальная обработка результата при работе с цепочкой

	fmt.Println("publish result")
	fmt.Println(result)

	if 

	body, err := json.Marshal(protocol.NewCeleryResult(s.SUCCESS, result, "", task.UUID()))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = b.Publisher.PublishWithContext(
		ctx,
		"",
		task.ReplyTo(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: task.CorrelationID(),
			Body:          body,
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

// TODO: Наверное можно вынести в один метод publish, но пока что пусть будет так
func (b *RabbitMQ) PublishException(result any, task interf.Tasks, trace string) error {
	body, err := json.Marshal(protocol.NewCeleryResult(s.FAILURE, result, trace, task.UUID()))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = b.Publisher.PublishWithContext(
		ctx,
		"",
		task.ReplyTo(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: task.CorrelationID(),
			Body:          body,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
