package config

type CeleryConfig struct {
	Broker BrokerSettings
	//Backend BackendSettings
	Worker WorkerSettings
	Queues []string
}

type BrokerSettings struct {
	Qos Qos
}

type Qos struct {
	PrefetchCount int
	PrefetchSize  int
	Global        bool
}

type WorkerSettings struct {
	WorkerConcurrency int
}
