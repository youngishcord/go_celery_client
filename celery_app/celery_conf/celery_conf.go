package celery_conf

// Надо посмотреть какие настройки я могу перенести
type CeleryConf struct {
	Broker  BrokerSettings
	Backend BackendSettings
	Worker  WorkerSettings
	Queues  []string
}

// Connection Не круто, что у меня тут пароль и логин но пока что так
type Connection struct {
	Host string
	Port string
	User string
	Pass string
}
