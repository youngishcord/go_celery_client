package celery_conf

import (
	"celery_client/celery_app/core/dto"
)

// Надо посмотреть какие настройки я могу перенести
type CeleryConf struct {
	Broker  dto.BrokerDto
	Backend dto.BackendDto
	Queues  []string
}
