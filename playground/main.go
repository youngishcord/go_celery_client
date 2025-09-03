package main

import "fmt"

type BaseMessages interface{}

type BaseTasks interface {
	ParseMessage([]any) (BaseMessages, error)
	RunTask(BaseMessages) ([]any, error)
}

type BaseTask struct {
	name string
}

func NewBaseTask(name string) BaseTask {
	return BaseTask{name: name}
}

type App struct {
	TasksRegistry map[string]BaseTasks
}

func (a *App) RegisterTask(name string, task BaseTasks) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = task
	return nil
}

type AddMessage struct {
	X, Y float64
}

type AddTask struct {
	BaseTask
}

func (t *AddTask) ParseMessage(data []any) (BaseMessages, error) {
	mess := &AddMessage{}

	if x, ok := data[0].(float64); ok {
		mess.X = x
	} else {
		return nil, fmt.Errorf("НЕКОРРЕКТНОЕ СООБЩЕНИЕ", data)
	}

	if y, ok := data[1].(float64); ok {
		mess.X = y
	} else {
		return nil, fmt.Errorf("НЕКОРРЕКТНОЕ СООБЩЕНИЕ", data)
	}

	return mess, nil
}

func (t *AddTask) RunTask(message AddMessage) ([]any, error) {
	return []any{message.X + message.Y}, nil
}

func main() {
	app := &App{
		TasksRegistry: map[string]BaseTasks{},
	}

	// mess := AddMessage{
	// 	X: 1.2,
	// 	Y: 2.3,
	// }
	task := AddTask{
		BaseTask: BaseTask{
			name: "add",
		},
	}

	app.RegisterTask("add", task)

	// app.TasksRegistry["add"].
}
