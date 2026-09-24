package exceptions

const (
	module = "celery.exceptions"
)

var Exception BaseException = BaseException{
	ExceptionType:   "Exception",
	ExceptionModule: "builtins",
}
