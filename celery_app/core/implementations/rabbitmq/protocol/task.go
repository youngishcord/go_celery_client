package amqp_protocol

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Args   []any          `json:"args"`
	Kwargs map[string]any `json:"kwargs"`
	Embed
}

// Embed Тут оставлены any типы пока что, в дальнейшем можно изменить
type Embed struct {
	Callbacks any `json:"callbacks,omitempty"`
	Errbacks  any `json:"errbacks,omitempty"`
	Chain     any `json:"chain,omitempty"`
	Chord     any `json:"chord,omitempty"`
}

func ParseTask(jsonData []byte) (Task, error) {
	fmt.Println(jsonData)
	var data []any
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return Task{}, err
	}

	if len(data) < 3 {
		return Task{}, fmt.Errorf("invalid data format: expected 3 elements, got %d", len(data))
	}

	task := Task{
		Embed: Embed{
			Callbacks: nil,
			Errbacks:  nil,
			Chain:     nil,
			Chord:     nil,
		},
	}

	// Парсим args (первый элемент)
	if args, ok := data[0].([]any); ok {
		task.Args = args
	}

	// Парсим kwargs (второй элемент)
	if kwargs, ok := data[1].(map[string]any); ok {
		task.Kwargs = kwargs
	}

	// Парсим вспомогательные данные (третий элемент)
	if auxData, ok := data[2].(map[string]any); ok {
		if callbacks, exists := auxData["callbacks"]; exists {
			task.Callbacks = callbacks
		}
		if errbacks, exists := auxData["errbacks"]; exists {
			task.Errbacks = errbacks
		}
		if chain, exists := auxData["chain"]; exists {
			task.Chain = chain
		}
		if chord, exists := auxData["chord"]; exists {
			task.Chord = chord
		}
	}

	return task, nil
}
