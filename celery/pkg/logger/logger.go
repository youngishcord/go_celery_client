package logger

type Option func(*Logger)

func NewLogger(options ...Option) *Logger {
	var l = Logger{}

	for _, opt := range options {
		opt(&l)
	}

	return &l
}

func WithIsJson() Option {
	return func(l *Logger) {
		l.Debug("text")
	}
}
