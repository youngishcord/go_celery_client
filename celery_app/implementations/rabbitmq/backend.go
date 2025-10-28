package rabbit

import (
	r "celery_client/celery_app/message/result"
	"context"
	"encoding/json"
	"fmt"
	"time"

	s "celery_client/celery_app/core/dto"
	protocol "celery_client/celery_app/core/dto/protocol"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Отношение к интерфейсу backend при работе с RPC
func (b *Rabbit) ConsumeResult(taskID string) (<-chan r.CeleryResult, error) {
	// TODO: implement me
	panic("IMPLEMENT ME")
}

// Отношение к интерфейсу backend при работе с RPC
func (b *Rabbit) PublishResult(result any, task protocol.CeleryTask) error {
	// TODO: тут стоит задуматься над тем что будет, если во второй операции выпадет ошибка, а
	//  первая уже будет выполнена
	// TODO: Подтверждение результата должно быть другим. Я возвращаю тут результат
	//  сразу в нужное место используя amqp либу и только после успешной публикации делаю Ack

	// TODO: Тут также нужна специальная обработка результата при работе с цепочкой

	fmt.Println("publish result")
	fmt.Println(result)

	//if

	body, err := json.Marshal(protocol.NewCeleryResult(s.SUCCESS, result, "", task.Headers.Id))
	if err != nil {
		return err
	}

	// TODO: тут надо поправить работу с функцией обработки таймаута
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = b.Publisher.PublishWithContext(
		ctx,
		"",
		task.Properties.ReplyTo.String(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: task.Properties.CorrelationID.String(),
			Body:          body,
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// FIXME: Тут может случиться так, что выполнение не подтвердится, что тогда делать, я хз
	err = b.Consumer.Ack(task.Properties.DeliveryTag, false)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Наверное можно вынести в один метод publish, но пока что пусть будет так
func (b *Rabbit) PublishException(result any, task protocol.CeleryTask, trace string) error {
	body, err := json.Marshal(protocol.NewCeleryResult(s.FAILURE, result, trace, task.Headers.Id))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = b.Publisher.PublishWithContext(
		ctx,
		"",
		task.Properties.ReplyTo.String(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: task.Properties.CorrelationID.String(),
			Body:          body,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	err = b.Consumer.Ack(task.Properties.DeliveryTag, false)
	if err != nil {
		return err
	}

	return nil
}
