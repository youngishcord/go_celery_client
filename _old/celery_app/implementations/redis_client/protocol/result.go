package redis_protocol

import "github.com/google/uuid"
import s "celery_client/celery_app/core/dto"
import "time"

type RedisCeleryResults struct {
	Status   s.Status  `json:"status"`
	Result   any       `json:"result"`
	Trace    string    `json:"traceback"`
	TaskID   uuid.UUID `json:"task_id"`
	Children []any     `json:"children"` // TODO: тут должен быть формат определенной формы, наверное.
	DateDone time.Time `json:"date_done"`
}

func NewCeleryResult(status s.Status, result any, trace string, taskID uuid.UUID) RedisCeleryResults {
	return RedisCeleryResults{
		Status:   status,
		Result:   result,
		Trace:    trace,
		TaskID:   taskID,
		DateDone: time.Now(),
		Children: make([]any, 0),
	}
}
