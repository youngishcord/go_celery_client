package exceptions

type ExceptionInfo struct {
	ExceptionType    string   `json:"exc_type"`
	ExceptionMessage []string `json:"exc_message"` // TODO: Точно ли тут всегда будут строки.
	ExceptionModule  string   `json:"exc_module"`
}

func NewExceptionInfo(err error) *ExceptionInfo {
	ex := EXCEPTION

	errMessage := []string{err.Error()}

	//for key, value := range goToPyException {
	//
	//}

	return &ExceptionInfo{
		ExceptionType:    ex.String(),
		ExceptionMessage: errMessage,
		ExceptionModule:  "", // TODO: Нужны какие то кастомные записи, или builtin??
	}
}
