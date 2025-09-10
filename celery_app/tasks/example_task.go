package base_tasks

import (
	protocol "celery_client/celery_app/core/implementations/rabbitmq/protocol"
	"fmt"
)

type AddTask struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
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

// Только переменные
func NewAddTask(message []byte) (BaseTasks, error) {
	parseTask, err := protocol.ParseTask(message)
	if err != nil {
		return nil, err
	}
	fmt.Println(parseTask)
	task := AddTask{
		X:        parseTask.Args[0].(float64),
		Y:        parseTask.Args[1].(float64),
		BaseTask: BaseTask{name: "name"},
	}
	//err = json.Unmarshal(message, &task)
	//if err != nil {
	//	return nil, err
	//}

	return &task, nil
	//name, ok := message["name"].(string)
	//if !ok {
	//	panic("NO NAME ERROR")
	//}
	//
	//task := AddTask{
	//	BaseTask: NewBaseTask(name),
	//}
	//
	//if x, ok := message["x"]; ok {
	//	task.X = x.(float64)
	//}
	//if y, ok := message["y"]; ok {
	//	task.Y = y.(float64)
	//}
	//
	//return task, nil
}
