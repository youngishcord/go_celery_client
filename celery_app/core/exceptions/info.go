package exceptions

type ExceptionInfo struct {
	ExceptionType    string   `json:"exc_type"`
	ExceptionMessage []string `json:"exc_message"` // TODO: Точно ли тут всегда будут строки.
	ExceptionModule  string   `json:"exc_module"`
}

func NewExceptionInfo(excType string, excMessage []string, excModule string) *ExceptionInfo {
	return &ExceptionInfo{
		ExceptionType:    excType,
		ExceptionMessage: excMessage,
		ExceptionModule:  excModule, // TODO: Нужны какие то кастомные записи, или builtin??
	}
}
