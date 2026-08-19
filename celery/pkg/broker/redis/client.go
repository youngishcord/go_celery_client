package broker

type Client struct {
}

func NewClient(config Config) (*Client, error) {
	return &Client{}, nil
}
