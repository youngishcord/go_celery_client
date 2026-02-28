package celery_app

import (
	"fmt"
	"go_celery_client/celery/internal/task"
	"go_celery_client/celery/protocol"
)

func (a *CeleryApp) RegisterTask(name string, constructor func(task *protocol.CeleryTask) (task.Task, error)) error {
	if _, ok := a.taskRegistry[name]; ok {
		return fmt.Errorf("NAME_ALREADY_RESERVED")
	}
	a.taskRegistry[name] = constructor
	return nil
}
