package redis

import celery "celery_client/celery_app/core/message/result"
import interf "celery_client/celery_app/core/interfaces"

func (b *Redis) PublishResult(result celery.CeleryResult, task interf.Tasks) error {
	panic("implement me")
}
func (b *Redis) ConsumeResult(taskID string) (<-chan celery.CeleryResult, error) {
	panic("implement me")
}
