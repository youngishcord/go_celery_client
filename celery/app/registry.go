package app

import (
	"fmt"
	"go_celery_client/celery/protocol"
	"go_celery_client/celery/task"
)

// RegisterTask добавление задачи в реджестри
func (a *CeleryApp) RegisterTask(name string, constructor func(task *protocol.CeleryTask) (task.Task, error)) error {
	if _, ok := a.taskRegistry[name]; ok {
		return fmt.Errorf("NAME_ALREADY_RESERVED")
	}
	a.taskRegistry[name] = constructor
	return nil
}

// RegisterTasks регистрирует несколько задач в реджестри
func (a *CeleryApp) RegisterTasks(tasks map[string]func(task *protocol.CeleryTask) (task.Task, error)) error {
	for name, constructor := range tasks {
		if _, ok := a.taskRegistry[name]; ok {
			return fmt.Errorf("NAME_ALREADY_RESERVED")
		}
		a.taskRegistry[name] = constructor
	}
	return nil
}
