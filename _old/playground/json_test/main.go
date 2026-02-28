package main

import (
	"encoding/json"
	"fmt"
)

// Основная структура для представления данных
type TaskData struct {
	Args      []interface{}          `json:"args"`
	Kwargs    map[string]interface{} `json:"kwargs"`
	Callbacks interface{}            `json:"callbacks"`
	Errbacks  interface{}            `json:"errbacks"`
	Chain     interface{}            `json:"chain"`
	Chord     interface{}            `json:"chord"`
}

// Альтернативный вариант с более строгой типизацией
type StrictTaskData struct {
	Args      []int                  `json:"args"` // если известно, что args всегда числа
	Kwargs    map[string]interface{} `json:"kwargs"`
	Callbacks *[]interface{}         `json:"callbacks"`
	Errbacks  *[]interface{}         `json:"errbacks"`
	Chain     *[]interface{}         `json:"chain"`
	Chord     *[]interface{}         `json:"chord"`
}

// Функция для парсинга JSON
func ParseTaskData(jsonData []byte) (*TaskData, error) {
	var data []interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	if len(data) < 3 {
		return nil, fmt.Errorf("invalid data format: expected 3 elements, got %d", len(data))
	}

	task := &TaskData{
		Callbacks: nil,
		Errbacks:  nil,
		Chain:     nil,
		Chord:     nil,
	}

	// Парсим args (первый элемент)
	if args, ok := data[0].([]interface{}); ok {
		task.Args = args
	}

	// Парсим kwargs (второй элемент)
	if kwargs, ok := data[1].(map[string]interface{}); ok {
		task.Kwargs = kwargs
	}

	// Парсим вспомогательные данные (третий элемент)
	if auxData, ok := data[2].(map[string]interface{}); ok {
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

func main() {
	// Пример JSON данных
	jsonBytes := []byte(`[[1, 2], {}, {
        "callbacks": null,
        "errbacks": null,
        "chain": null,
        "chord": null
    }]`)

	task, err := ParseTaskData(jsonBytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Args: %v\n", task.Args)
	fmt.Printf("Kwargs: %v\n", task.Kwargs)
	fmt.Printf("Callbacks: %v\n", task.Callbacks)
	fmt.Printf("Errbacks: %v\n", task.Errbacks)
}
