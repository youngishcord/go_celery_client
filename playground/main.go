package main

import (
	"fmt"
	"sync"
)

type BaseMessages interface {
	Tmp()
}

type BaseTasks interface {
	Run() (any, error)
	Message() (any, error)
}

type TaskWrapper struct {
	RunTask      func(BaseMessages) ([]any, error)
	ParseMessage func([]any)
}

type BaseTask struct {
	name string
}

func NewBaseTask(name string) BaseTask {
	return BaseTask{name: name}
}

type App struct {
	TasksRegistry map[string]BaseTasks
	resultCh      chan any
}

func (a *App) RegisterTask(name string, task BaseTasks) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = task
	return nil
}

type AddTask struct {
	X, Y float64
	BaseTask
}

func (t *AddTask) Complete() {
	//TODO implement me
	panic("implement me")
}

func (t *AddTask) Message() (any, error) {
	// Похуй
	return 1, nil
}

func (t *AddTask) Run() (any, error) {
	if t == nil {
		panic("хуй")
	}
	fmt.Println("this is add task")
	return t.X + t.Y, nil
}

func NewAddTask(x, y float64) AddTask {
	return AddTask{
		X:        x,
		Y:        y,
		BaseTask: NewBaseTask("add_task"),
	}
}

func main() {
	app := &App{
		TasksRegistry: map[string]BaseTasks{},
		resultCh:      make(chan any, 20),
	}

	// mess := AddMessage{
	// 	X: 1.2,
	// 	Y: 2.3,
	// }
	task := NewAddTask(1.2, 2.4)

	err := app.RegisterTask("add", &task)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := app.TasksRegistry["add"].Run()
		if err != nil {
			return
		}
		//fmt.Println(res)
		app.resultCh <- res
	}()

	select {
	case res := <-app.resultCh:
		fmt.Println(res)
	}

	wg.Wait()
	// app.TasksRegistry["add"].
}
