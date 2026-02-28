package base_tasks

import (
	app "celery_client/celery_app/app"
	protocol "celery_client/celery_app/core/dto/protocol"

	_ "github.com/google/uuid"
)

type AddTask struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	protocol.CeleryTask
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

// Возвращаемый тип должен уподоблять интерфейсу задачи.
// В конструкторе необходимо парсить аргументы в структуру для дальнейшей работы
// и приводить типы. Наличие и остутствие переменных, равно как и их
// последовательность остается на разработчике.
func NewAddTask(rawTask protocol.CeleryTask) (app.Task, error) {
	args := rawTask.Body.Args
	task := AddTask{
		X:          args[0].(float64),
		Y:          args[1].(float64),
		CeleryTask: rawTask,
	}

	return &task, nil
}
