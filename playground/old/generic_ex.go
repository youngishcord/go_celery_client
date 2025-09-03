package main

import "fmt"

// Обобщённый интерфейс для задачи
type Task[M any] interface {
	ParseMessage([]any) (M, error)
	RunTask(M) ([]any, error)
}

// Реестр задач
type App struct {
	TasksRegistry map[string]any
}

func NewApp() *App {
	return &App{TasksRegistry: make(map[string]any)}
}

// Регистрация задачи
func RegisterTask[M any](a *App, name string, task Task[M]) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = task
	return nil
}

// Получение задачи с конкретным типом
func GetTask[M any](a *App, name string) (Task[M], error) {
	if task, ok := a.TasksRegistry[name]; ok {
		if typedTask, ok := task.(Task[M]); ok {
			return typedTask, nil
		}
		return nil, fmt.Errorf("ТИП СООБЩЕНИЯ НЕ СОВПАДАЕТ ДЛЯ ЗАДАЧИ %q", name)
	}
	return nil, fmt.Errorf("ЗАДАЧА %q НЕ НАЙДЕНА", name)
}

// ==== Конкретная задача ====

// Сообщение
type AddMessage struct {
	X, Y float64
}

// Задача
type AddTask struct{}

func (t *AddTask) ParseMessage(data []any) (AddMessage, error) {
	if len(data) < 2 {
		return AddMessage{}, fmt.Errorf("НЕ ХВАТАЕТ ДАННЫХ: %v", data)
	}

	x, ok1 := data[0].(float64)
	y, ok2 := data[1].(float64)
	if !ok1 || !ok2 {
		return AddMessage{}, fmt.Errorf("НЕКОРРЕКТНОЕ СООБЩЕНИЕ: %v", data)
	}

	return AddMessage{X: x, Y: y}, nil
}

func (t *AddTask) RunTask(m AddMessage) ([]any, error) {
	return []any{m.X + m.Y}, nil
}

// ==== main ====

func main() {
	app := NewApp()

	addTask := &AddTask{}
	if err := RegisterTask(app, "add", addTask); err != nil {
		panic(err)
	}

	task, err := GetTask[AddMessage](app, "add")
	if err != nil {
		panic(err)
	}

	msg, err := task.ParseMessage([]any{1.2, 2.3})
	if err != nil {
		panic(err)
	}

	result, err := task.RunTask(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Результат:", result)
}

/*
type BaseMessages interface {
	Tmp()
}

// Generic версия
type BaseTasks[M BaseMessages] interface {
	RunTask(M) ([]any, error)
	ParseMessage([]any) (M, error)
}

// Обёртка для хранения разных задач в одном map
type TaskWrapper struct {
	RunTask      func(BaseMessages) ([]any, error)
	ParseMessage func([]any) (BaseMessages, error)
}
tasks := map[string]TaskWrapper{
	"add": {
		RunTask: func(msg BaseMessages) ([]any, error) {
			m := msg.(*AddMessage)
			return (&Add{}).RunTask(m)
		},
		ParseMessage: func(data []any) (BaseMessages, error) {
			return (&Add{}).ParseMessage(data)
		},
	},
}

*/
