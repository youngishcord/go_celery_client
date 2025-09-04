package base_tasks

type BaseTasks interface {
	Run() (any, error)
	Message() (any, error)
}

type BaseTask struct {
	name string
}

func NewBaseTask(name string) BaseTask {
	return BaseTask{name: name}
}
