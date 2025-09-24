package exceptions

import "errors"
import e "celery_client/celery_app/core/errors"

var Exception BaseException = BaseException{
	ExceptionType:   "Exception",
	ExceptionModule: "builtins",
}

var goToPyException = map[string]BaseException{
	e.NotRegistered.Error(): BaseException{
		ExceptionType:   "NotRegistered",
		ExceptionModule: "celery.exceptions",
	},
}

func RegisterNewExceptions(ex map[string]BaseException) error {
	for key, value := range ex {
		if _, ok := goToPyException[key]; ok {
			return errors.New(key + " already exists")
		}
		goToPyException[key] = value
	}
	return nil
}

func GetException(err error, exceptionMessage []string) *ExceptionInfo {
	var exception BaseException = Exception
	if exc, ok := goToPyException[err.Error()]; ok {
		exception = exc
	}
	return NewExceptionInfo(
		exception.ExceptionType,
		exceptionMessage,
		exception.ExceptionModule,
	)
}
