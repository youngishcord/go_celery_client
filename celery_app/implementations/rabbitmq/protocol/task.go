package amqp_protocol

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Args   []any          `json:"args"`
	Kwargs map[string]any `json:"kwargs"`
	Emb    Embed          `json:"embed"`
}

type Options struct {
	Queue   string `json:"queue"`
	TaskId  string `json:"task_id"`
	ReplyTo string `json:"reply_to"`
}

type Chain struct {
	TaskName    string         `json:"task"`
	Args        []any          `json:"args"`
	Kwargs      map[string]any `json:"kwargs"`
	Opt         Options        `json:"options"`
	SubtaskType any            `json:"subtask_type"`
	Immutable   bool           `json:"immutable"`
}

// Embed Тут оставлены any типы пока что, в дальнейшем можно изменить
type Embed struct {
	Callbacks any     `json:"callbacks,omitempty"`
	Errbacks  any     `json:"errbacks,omitempty"`
	Chain     []Chain `json:"chain,omitempty"`
	Chord     any     `json:"chord,omitempty"`
}

func ParseTask(jsonData []byte) (Task, error) {
	fmt.Println(jsonData)
	var data []json.RawMessage
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return Task{}, err
	}

	if len(data) < 3 {
		return Task{}, fmt.Errorf("invalid data format: expected 3 elements, got %d", len(data))
	}

	task := Task{}

	// Парсим args (первый элемент)
	if err := json.Unmarshal(data[0], &task.Args); err != nil {
		panic("NO ARGS")
	}
	// if args, ok := data[0].([]any); ok {
	// 	task.Args = args
	// }

	// Парсим kwargs (второй элемент)
	if err := json.Unmarshal(data[1], &task.Kwargs); err != nil {
		panic("NO KWARGS")
	}
	// if kwargs, ok := data[1]; ok {
	// 	task.Kwargs = kwargs
	// }

	// emb := Embed{}
	// Парсим вспомогательные данные (третий элемент)
	err = json.Unmarshal(data[2], &task.Emb)
	if err != nil {
		panic("TASK CHAIN ERROR")
	}
	// } else {
	// 	task.Emb = emb
	// }
	// if auxData, ok := data[2].(map[string][]byte); ok {
	// if callbacks, exists := auxData["callbacks"]; exists {
	// 	task.Callbacks = callbacks
	// }
	// if errbacks, exists := auxData["errbacks"]; exists {
	// 	task.Errbacks = errbacks
	// }
	// if chain, exists := auxData["chain"]; exists {
	// 	ch := Chain{}
	// 	err := json.Unmarshal(chain, ch)
	// 	if err != nil {
	// 		fmt.Errorf("CHAIN TASK ERROR")
	// 	}

	// }
	// if chord, exists := auxData["chord"]; exists {
	// 	task.Chord = chord
	// }
	// }

	fmt.Println("FORMATED TASK ")
	fmt.Println(task)

	return task, nil
}
