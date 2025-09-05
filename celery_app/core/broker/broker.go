package broker

type Broker interface {
	Connect(queues []string) error
	TaskChannel() *chan 
}
