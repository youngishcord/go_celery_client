package exceptions

import "errors"

type BaseException string

// Хз надо ли оно мне
const (
	EXCEPTION  BaseException = "Exception"
	VALUEERROR BaseException = "ValueError"
)

func (exception BaseException) String() string {
	return string(exception)
}

var goToPyException = map[string]string{
	"invalid argument": "ValueError",
	"not found":        "KeyError",
	"permission":       "PermissionError",
}

func RegisterNewExceptions(ex map[string]string) error {
	for key, value := range ex {
		if _, ok := goToPyException[key]; ok {
			return errors.New(goToPyException[key] + " already exists")
		}
		goToPyException[key] = value
	}
	return nil
}
