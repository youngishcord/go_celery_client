package dto

type BackendDto struct {
	BackendType    string
	ConnectionData Connection
}

type BrokerDto struct {
	BrokerType     string
	ConnectionData Connection
}

// Connection Не круто, что у меня тут пароль и логин но пока что так
type Connection struct {
	Host string
	Port string
	User string
	Pass string
}
