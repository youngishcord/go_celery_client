package celery_conf

const (
	SoloConcurrency int = 1
)

type WorkerSettings struct {
	WorkerConcurrency int
}
