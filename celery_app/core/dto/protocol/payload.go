package protocol

import (
	"encoding/json"
	"fmt"
)

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
	Chord     any     `json:"chord,omitempty"` // TODO: доработать
}

type Body struct {
	Args   []any          `json:"args"`
	Kwargs map[string]any `json:"kwargs"`
	Emb    Embed          `json:"embed"`
}

func ParsePayload(jsonData []byte) (Body, error) {
	fmt.Println(jsonData)
	var data []json.RawMessage
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return Body{}, err
	}

	if len(data) < 3 {
		return Body{}, fmt.Errorf("invalid data format: expected 3 elements, got %d", len(data))
	}

	task := Body{}

	// Парсим args (первый элемент)
	if err := json.Unmarshal(data[0], &task.Args); err != nil {
		panic("NO ARGS")
	}

	// Парсим kwargs (второй элемент)
	if err := json.Unmarshal(data[1], &task.Kwargs); err != nil {
		panic("NO KWARGS")
	}

	// Парсим вспомогательные данные (третий элемент)
	err = json.Unmarshal(data[2], &task.Emb)
	if err != nil {
		panic("TASK CHAIN ERROR")
	}

	fmt.Println("FORMATED TASK ")
	fmt.Println(task)

	return task, nil
}
