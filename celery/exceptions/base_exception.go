package exceptions

type BaseException struct {
	ExceptionType   string `json:"exc_type"`
	ExceptionModule string `json:"exc_module"`
}
