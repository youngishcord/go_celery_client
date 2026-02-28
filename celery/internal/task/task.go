package task

import "context"

// Task отвечает за работу с задачами приложения
type Task interface {
	Run(ctx context.Context) (any, error)
	Name() string
}
