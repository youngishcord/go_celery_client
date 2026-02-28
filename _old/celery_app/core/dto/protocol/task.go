package protocol

// Универсальный формат задачи, которую можно таскать по программе
type CeleryTask struct {
	ContentEncoding string     `json:"content-encoding"`
	ContentType     string     `json:"content-type"`
	Body            Body       `json:"body"`
	Headers         Header     `json:"headers"`
	Properties      Properties `json:"properties"`
}
