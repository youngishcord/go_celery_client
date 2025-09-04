package amqp_models

import (
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Header struct {
	Lang     string    `json:"lang"`
	Task     string    `json:"task"`
	Id       uuid.UUID `json:"id"`
	RootId   uuid.UUID `json:"root_id,omitempty"`
	ParentId uuid.UUID `json:"parent_id,omitempty"`
	GroupId  uuid.UUID `json:"group,omitempty"`

	// optional
	Meth                string     `json:"meth,omitempty"`
	Shadow              string     `json:"shadow,omitempty"`
	ETA                 *time.Time `json:"eta,omitempty"`
	Expires             *time.Time `json:"expires,omitempty"`
	Retries             int        `json:"retries,omitempty"`
	TimeLimit           *TimeLimit `json:"timelimit,omitempty"`
	ArgsRepr            string     `json:"argsrepr,omitempty"`
	KwargsRepr          string     `json:"kwargsrepr,omitempty"`
	Origin              string     `json:"origin,omitempty"`
	ReplacedTaskNesting int        `json:"replaced_task_nesting,omitempty"`
}

type TimeLimit struct {
	Soft time.Duration `json:"soft"`
	Hard time.Duration `json:"hard"`
}

func MakeHeaderFromMap(data amqp.Table) (*Header, error) {

	header := &Header{}

	if lang, ok := data["lang"].(string); ok {
		header.Lang = lang
	} else {
		header.Lang = "py"
	}

	if task, ok := data["task"].(string); ok {
		header.Task = task
	}

	if idStr, ok := data["id"].(string); ok {
		if id, err := uuid.Parse(idStr); err == nil {
			header.Id = id
		}
	}

	if rootIdStr, ok := data["root_id"].(string); ok {
		if id, err := uuid.Parse(rootIdStr); err == nil {
			header.RootId = id
		}
	}

	if parentIdStr, ok := data["parent_id"].(string); ok {
		if id, err := uuid.Parse(parentIdStr); err == nil {
			header.ParentId = id
		}
	}

	if grIdStr, ok := data["group"].(string); ok {
		if id, err := uuid.Parse(grIdStr); err == nil {
			header.GroupId = id
		}
	}

	if meth, ok := data["meth"].(string); ok {
		header.Meth = meth
	}

	if shadow, ok := data["shadow"].(string); ok {
		header.Shadow = shadow
	}

	if etaStr, ok := data["eta"].(string); ok {
		if eta, err := time.Parse(time.RFC3339, etaStr); err == nil {
			header.ETA = &eta
		}
	}

	if expiresStr, ok := data["expires"].(string); ok {
		if expires, err := time.Parse(time.RFC3339, expiresStr); err == nil {
			header.Expires = &expires
		}
	}

	if retries, ok := data["retries"].(float64); ok {
		header.Retries = int(retries)
	}

	if retries, ok := data["retries"].(int); ok {
		header.Retries = retries
	}

	if timelimit, ok := data["timelimit"].([]interface{}); ok && len(timelimit) == 2 {
		if soft, ok := timelimit[0].(float64); ok {
			if hard, ok := timelimit[1].(float64); ok {
				header.TimeLimit = &TimeLimit{
					Soft: time.Duration(soft) * time.Second,
					Hard: time.Duration(hard) * time.Second,
				}
			}
		}
	}

	if argsrepr, ok := data["argsrepr"].(string); ok {
		header.ArgsRepr = argsrepr
	}

	if kwargsrepr, ok := data["kwargsrepr"].(string); ok {
		header.KwargsRepr = kwargsrepr
	}

	if origin, ok := data["origin"].(string); ok {
		header.Origin = origin
	}

	if nesting, ok := data["replaced_task_nesting"].(float64); ok {
		header.ReplacedTaskNesting = int(nesting)
	}

	if nesting, ok := data["replaced_task_nesting"].(int); ok {
		header.ReplacedTaskNesting = nesting
	}

	return header, nil
}
