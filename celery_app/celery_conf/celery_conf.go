package celery_conf

import dto "celery_client/celery_app/dto"

// Надо посмотреть какие настройки я могу перенести
type CeleryConf struct {
	Broker  dto.BrokerDto
	Backend dto.BackendDto
	Queues  []string
}
