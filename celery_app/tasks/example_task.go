package base_tasks

type AddTask struct {
	X, Y float64
	BaseTask
}

func (t *AddTask) Message() (any, error) {
	// Похуй
	return 1, nil
}

func (t *AddTask) Run() (any, error) {
	if t == nil {
		panic("хуй")
	}
	return t.X + t.Y, nil
}

func NewAddTask(x, y float64) AddTask {
	return AddTask{
		X:        x,
		Y:        y,
		BaseTask: NewBaseTask("add_task"),
	}
}
