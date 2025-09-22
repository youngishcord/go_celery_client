package redis_client

import (
	s "celery_client/celery_app/core/dto"
	interf "celery_client/celery_app/core/interfaces"
	protocol "celery_client/celery_app/implementations/redis_client/protocol"
	celery "celery_client/celery_app/message/result"
	"context"
	"encoding/json"
	"time"
)

// TODO: ttl должен быть настраиваемый
func (b *RedisClient) PublishResult(result any, task interf.BaseTasks) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	body, err := json.Marshal(protocol.NewCeleryResult(s.SUCCESS, result, "", task.UUID()))
	if err != nil {
		return err
	}
	status := b.conn.Set(ctx, "celery-task-meta-"+task.CorrelationID(), body, b.ttl)
	if status != nil {
		println(status)
	}
	return nil
}

func (b *RedisClient) ConsumeResult(taskID string) (<-chan celery.CeleryResult, error) {
	panic("implement me")
}
