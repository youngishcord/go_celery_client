package celery_conf

import dto "celery_client/celery_app/dto"

// Надо посмотреть какие настройки я могу перенести
type CeleryConf struct {
	Broker  dto.Connection
	Backend dto.Connection
	Queues  []string
}
