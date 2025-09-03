package worker

import "fmt"

// type BaseTask func([]interface{}) (interface{}, error)
// Интерфейс для описания задач, которые будут храниться в regestry
type BaseTasks interface {
	GetName() string
	RunTask() (any, error)
	// Какие методы необходимы:
	// 1. Передача параметров динамически
	// 2. Формирование ответа
	// 3. Формирование задачи с параметрами от функции
	// 4.
	// 5.
}

// Базовая задача, в которой будут храниться типовые поля и методы
type BaseTask struct {
	name string
}

// Типовой метод получения имени задачи
func (t *BaseTask) GetName() string {
	return t.name
}

func NewBaseTask(name string) BaseTask {
	return BaseTask{name: name}
}

type Worker struct {
	TasksRegistry map[string]BaseTasks
}

type Add struct {
	x float64
	y float64

	// Подключение полей типовой задачи в структуру своей задачи
	BaseTask
}

func (a *Add) RunTask() (any, error) {
	return a.x + a.y, nil
}

func TestF() {
	reg := Worker{
		TasksRegistry: map[string]BaseTasks{},
	}
	reg.TasksRegistry["test"] = &Add{
		x:        1,
		y:        2,
		BaseTask: NewBaseTask("test"),
	}

	fmt.Println(reg)

	fmt.Println(reg.TasksRegistry["test"].RunTask())
	fmt.Println(reg.TasksRegistry["test"].GetName())
}
