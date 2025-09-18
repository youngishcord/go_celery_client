package redis_client

import (
	celery "celery_client/celery_app/message/result"
	"context"
	"encoding/json"
	"time"
)
import interf "celery_client/celery_app/core/interfaces"

func (b *RedisClient) PublishResult(result celery.CeleryResult, baseTasks interf.BaseTasks) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	body, err := json.Marshal(result)
	if err != nil {
		return err
	}
	status := b.conn.Set(ctx, "celery-task-meta-"+baseTasks.CorrelationID(), body, b.ttl)
	if status != nil {
		println(status)
	}
	return nil
}

func (b *RedisClient) ConsumeResult(taskID string) (<-chan celery.CeleryResult, error) {
	panic("implement me")
}
