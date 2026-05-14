package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Option func(*Logger)

func NewLogger(options ...Option) *Logger {

	rotator := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    1, // MB
		MaxBackups: 2,
		MaxAge:     7, // days
		Compress:   false,
	}

	multi := io.MultiWriter(os.Stdout, rotator)

	handler := slog.NewJSONHandler(multi, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})

	return slog.New(handler)

}

func WithIsJson() Option {
	return func(l *Logger) {
		l.Debug("text")
	}
}

func WithOutput(output io.Writer) Option {
	return func(l *Logger) {

	}
}

func WithRotationWriter(cfg RotationConfig) Option {
	return func(l *Logger) {
		NewRotatingWriter(cfg)
	}
}
