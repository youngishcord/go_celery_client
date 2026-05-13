package celery_app

import (
	"go_celery_client/celery/internal/worker"
	"log"
)

// Start запускает пул воркеров
func (a *CeleryApp) Start() error {
	err := a.broker.Start([]string{"test"})
	if err != nil {
		return err
	}

	for i := 0; i < a.conf.Worker.WorkerConcurrency; i++ {
		w, err := worker.NewCeleryWorker(i, a, a.workerPool.closeCh)
		if err != nil {
			return err
		}

		go func() {
			err := w.Start()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	return nil
}
