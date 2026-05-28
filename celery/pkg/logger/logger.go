package logger

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultLevel             = LevelInfo
	defaultStdOut            = true
	defaultAddSource         = true
	defaultSetDefault        = false
	defaultToFile            = false
	defaultLogFile           = "celery.log"
	defaultLogFileMaxSizeMB  = 10
	defaultLogFileMaxBackups = 5
	defaultLogFileMaxAgeDays = 14
	defaultLogFileCompress   = true
)

type LoggerOptions struct {
	Level             Level
	StdOut            bool
	AddSource         bool
	SetDefault        bool
	ToFile            bool
	LogFilePath       string
	LogFileMaxSizeMB  int
	LogFileMaxBackups int
	LogFileMaxAgeDays int
	LogFileCompress   bool
	ProcessName       string
}

type Option func(*LoggerOptions)

func NewLogger(options ...Option) *Logger {
	cfg := &LoggerOptions{
		Level:             defaultLevel,
		StdOut:            defaultStdOut,
		AddSource:         defaultAddSource,
		SetDefault:        defaultSetDefault,
		ToFile:            defaultToFile,
		LogFilePath:       defaultLogFile,
		LogFileMaxSizeMB:  defaultLogFileMaxSizeMB,
		LogFileMaxBackups: defaultLogFileMaxBackups,
		LogFileMaxAgeDays: defaultLogFileMaxAgeDays,
		LogFileCompress:   defaultLogFileCompress,
	}

	for _, opt := range options {
		opt(cfg)
	}

	var w []io.Writer

	if cfg.ToFile {
		rotation := &lumberjack.Logger{
			Filename:   cfg.LogFilePath,
			MaxSize:    cfg.LogFileMaxSizeMB, // MB
			MaxBackups: cfg.LogFileMaxBackups,
			MaxAge:     cfg.LogFileMaxAgeDays, // days
			Compress:   cfg.LogFileCompress,
		}

		w = append(w, rotation)
	}

	if cfg.StdOut {
		w = append(w, os.Stdout)
	}

	writer := io.MultiWriter(w...)

	//handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
	//	Level:     slog.LevelInfo,
	//	AddSource: true,
	//})

	h := NewCeleryHandler(writer, HandlerOptions{
		Level:       cfg.Level,
		AddSource:   cfg.AddSource,
		ProcessName: cfg.ProcessName,
	})

	logger := New(h)

	if cfg.SetDefault {
		SetDefault(logger)
	}

	return logger
}

func NoStdOutput(output io.Writer) Option {
	return func(l *LoggerOptions) {
		l.StdOut = false
	}
}

func WithRotationWriter(cfg RotationConfig) Option {
	return func(l *LoggerOptions) {
		l.ToFile = true
	}
}

func WithSetDefault() Option {
	return func(l *LoggerOptions) {
		l.SetDefault = true
	}
}
