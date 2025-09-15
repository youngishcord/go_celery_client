package base_tasks

import (
	interf "celery_client/celery_app/core/interfaces"

	_ "github.com/google/uuid"
)

type AddTask struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	BaseTask
}

func (t *AddTask) ReplyTo() string {
	//TODO implement me
	panic("implement me")
}

// func (t *AddTask) Complete() {
// 	t.rawTask.Ack()
// }

func (t *AddTask) Message() (any, error) {
	// Похуй
	return 1, nil
}

func (t *AddTask) Run() (any, error) { // тут надо сделать сериализемый возвращаемый тип
	if t == nil {
		panic("хуй")
	}
	return t.X + t.Y, nil
}

// Только переменные
func NewAddTask(rawTask interf.Tasks) (BaseTasks, error) {
	//parseTask, err := protocol.ParseTask(message)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(parseTask)
	args := rawTask.Args()
	task := AddTask{
		X:        args[0].(float64),
		Y:        args[1].(float64),
		BaseTask: NewBaseTask(rawTask),
		// BaseTask: BaseTask{name: "name"},
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
