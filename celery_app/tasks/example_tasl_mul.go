package examples

import (
	app "celery_client/celery_app/app"
	protocol "celery_client/celery_app/core/dto/protocol"
	"context"
	"time"

	_ "github.com/google/uuid"
)

type Pow struct {
	X float64 `json:"x"`
	protocol.CeleryTask
}

func (t *Pow) Message() (any, error) {
	// Похуй
	return 1, nil
}

func (t *Pow) Run(ctx context.Context) (any, error) {
	time.Sleep(1 * time.Second)
	if t == nil {
		panic("хуй")
	}
	return t.X * t.X, nil
}

// Возвращаемый тип должен уподоблять интерфейсу задачи.
// В конструкторе необходимо парсить аргументы в структуру для дальнейшей работы
// и приводить типы. Наличие и остутствие переменных, равно как и их
// последовательность остается на разработчике.
func NewPowTask(rawTask protocol.CeleryTask) (app.Task, error) {
	args := rawTask.Body.Args
	task := AddTask{
		X:          args[0].(float64),
		CeleryTask: rawTask,
	}

	return &task, nil
}
