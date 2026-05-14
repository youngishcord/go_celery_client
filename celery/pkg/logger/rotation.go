package logger

import (
	"io"

	lj "gopkg.in/natefinch/lumberjack.v2"
)

type RotationConfig struct {
	Filename   string
	MaxSize    int // MB
	MaxAge     int // Days
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

func NewRotatingWriter(cfg RotationConfig) io.WriteCloser {
	return &lj.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
}
