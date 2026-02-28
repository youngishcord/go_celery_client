package broker

import "fmt"

type Config struct {
	Host string
	Port string
	User User
}

type User struct {
	Username string
	Password string
}

func (c *Config) Url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.User.Username, c.User.Password, c.Host, c.Port)
}
