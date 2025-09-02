package worker

// TaskFunc Сомнительно
type TaskFunc func([]interface{}) (interface{}, error)

type Worker struct {
	TasksRegistry map[string]TaskFunc
}
