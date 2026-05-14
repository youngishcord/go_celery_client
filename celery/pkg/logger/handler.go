package logger

import (
	"context"
	"io"
)

type customHandler interface {
	Enabled(context.Context, Level) bool
	Handle(context.Context, Record) error
	WithAttrs(attrs []Attr) Handler
	WithGroup(name string) Handler
}

type CeleryHandler struct {
	opts   HandlerOptions
	writer io.Writer
	attrs  []Attr
	groups []string
}

func (h *CeleryHandler) Enabled(ctx context.Context, l Level) bool {
	minLevel := LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}
func (h *CeleryHandler) Handle(ctx context.Context, r Record) error {
	fields := make(map[string]any)

	fields["timestamp"] = r.Time.Format("2006-01-02T15:04:05.000Z07:00")
	fields["level"] = r.Level.String()
	fields["message"] = r.Message

	for _, attr := range h.attrs {
		fields[attr.Key] = attr.Value
	}

	r.Attrs(func(attr Attr) bool {
		fields[attr.Key] = attr.Value.Any()
		return true
	})

	return nil
}
func (h *CeleryHandler) WithAttrs(attrs []Attr) Handler {
	return &CeleryHandler{
		opts:   HandlerOptions{},
		writer: nil,
		attrs:  nil,
		groups: nil,
	}
}
func (h *CeleryHandler) WithGroup(name string) Handler {
	return &CeleryHandler{
		opts:   HandlerOptions{},
		writer: nil,
		attrs:  nil,
		groups: nil,
	}
}
