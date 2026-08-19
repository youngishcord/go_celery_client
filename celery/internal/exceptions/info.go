package exceptions

type ExceptionInfo struct {
	ExceptionType    string         `json:"exc_type"`
	ExceptionMessage []string       `json:"exc_message"`
	ExceptionModule  string         `json:"exc_module"`
	Args             []any          `json:"args"`
	Kwargs           map[string]any `json:"kwargs"`
}

func NewExceptionInfo(excType string, excMessage []string, excModule string, args []any, kwargs map[string]any) *ExceptionInfo {
	return &ExceptionInfo{
		ExceptionType:    excType,
		ExceptionMessage: excMessage,
		ExceptionModule:  excModule,
		Args:             args,
		Kwargs:           kwargs,
	}
}
