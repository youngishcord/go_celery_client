package redis_client

import (
	s "celery_client/celery_app/core/dto"
	"celery_client/celery_app/core/dto/protocol"
	result "celery_client/celery_app/message/result"
	"context"
	"encoding/json"
	"time"
)

// TODO: ttl должен быть настраиваемый
func (b *RedisClient) PublishResult(result any, task protocol.CeleryTask) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	body, err := json.Marshal(protocol.NewCeleryResult(s.SUCCESS, result, "", task.Headers.Id))
	if err != nil {
		return err
	}
	status := b.conn.Set(ctx, "celery-task-meta-"+task.Properties.CorrelationID.String(), body, b.ttl)
	if status != nil {
		println(status)
	}
	return nil
}

func (b *RedisClient) ConsumeResult(taskID string) (<-chan result.CeleryResult, error) {
	panic("implement me")
}

func (b *RedisClient) PublishException(result any, baseTasks protocol.CeleryTask, trace string) error {
	//TODO implement me
	panic("implement me")
}
