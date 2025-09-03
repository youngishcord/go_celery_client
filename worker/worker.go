package worker

import (
	"fmt"
	"log"
)

type BaseMessages interface{}

type BaseMessage struct{}

// type BaseTask func([]interface{}) (interface{}, error)
// Интерфейс для описания задач, которые будут храниться в regestry
type BaseTasks[M BaseMessages] interface {
	GetName() string
	RunTask(M) ([]any, error)      // Запуск исполнения задач
	ParseMessage([]any) (M, error) // Формирование структуры сообщения из поступивших данных

	// Какие методы необходимы:
	// 1. Передача параметров динамически
	// 2. Формирование ответа
	// 3. Формирование задачи с параметрами от функции
	// 4.
	// 5.
}

type TaskWrapper struct {
	RunTask      func(BaseMessages) ([]any, error)
	ParseMessage func([]any) (BaseMessages, error)
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

type App struct {
	TasksRegistry map[string]BaseTasks
}

func (a *App) RegisterTask(name string, task BaseTasks) error {
	if _, ok := a.TasksRegistry[name]; ok {
		return fmt.Errorf("ЗАДАЧА С ТАКИМ ИМЕНЕМ УЖЕ ЗАРЕГИСТРИРОВАНА")
	}
	a.TasksRegistry[name] = {
		RunTask: func(msg BaseMessages) ([]any, error) {
			m := msg.(*AddMessage)
			return (&Add{}).RunTask(m)
		},
		ParseMessage: func(data []any) (BaseMessages, error) {
			return (&Add{}).ParseMessage(data)
		},
	},
	return nil
}

type AddMessage struct {
	x float64
	y float64
}

func (m *AddMessage) Tmp() {
	return
}

type Add struct {
	// Подключение полей типовой задачи в структуру своей задачи
	BaseTask
}

// Тут важно понимать что именно приходит в args и в kwargs
// Нужно правильно обрабатывать args и kwargs
func (t *Add) parseMessage(data []any) (BaseMessages, error) {
	mess := &AddMessage{}

	if x, ok := data[0].(float64); ok {
		mess.x = x
	} else {
		return nil, fmt.Errorf("НЕКОРРЕКТНОЕ СООБЩЕНИЕ", data)
	}

	if y, ok := data[1].(float64); ok {
		mess.y = y
	} else {
		return nil, fmt.Errorf("НЕКОРРЕКТНОЕ СООБЩЕНИЕ", data)
	}

	return mess, nil
}

func (t *Add) RunTask(message []any) ([]any, error) {
	mess := t.parseMessage(message)
	return []any{mess.x + mess.y}, nil
}

// type Mul struct {
// 	a float64
// 	b float64

// 	BaseTask
// }

// func (t *Mul) RunTask() ([]any, error) {
// 	return []any{t.a * t.b}, nil
// }

func TestF() {
	reg := App{
		TasksRegistry: map[string]TaskWrapper{},
	}

	err := reg.RegisterTask("add", &Add{
		BaseTask: NewBaseTask("add"),
	})

	if err != nil {
		log.Fatalln(err)
	}

	// reg.TasksRegistry["test"] = &Add{
	// 	x:        1,
	// 	y:        2,
	// 	BaseTask: NewBaseTask("test"),
	// }

	// reg.TasksRegistry["mul"] = &Mul{
	// 	a:        10,
	// 	b:        11,
	// 	BaseTask: NewBaseTask("mul"),
	// }

	fmt.Println(reg)

	// mess, err := reg.TasksRegistry["add"].ParseMessage([]any{1.0, 2.5})
	fmt.Println(reg.TasksRegistry["add"].RunTask([]any{1.0, 2.5}))
	fmt.Println(reg.TasksRegistry["add"].GetName())

	// fmt.Println(reg.TasksRegistry["mul"].RunTask())
	// fmt.Println(reg.TasksRegistry["mul"].GetName())
}
