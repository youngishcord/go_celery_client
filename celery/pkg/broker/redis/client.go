package broker

type Client struct {
}

type Config struct {
	User struct {
		Username string
		Password string
	}
}

func NewClient(config Config) (*Client, error) {
	return &Client{}, nil
}
