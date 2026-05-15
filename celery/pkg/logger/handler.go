package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

type HandlerOptions struct {
	Level       slog.Leveler
	AddSource   bool
	ProcessName string
}

type CeleryHandler struct {
	opts   HandlerOptions
	writer io.Writer
	attrs  []Attr
	groups []string
	mu     sync.Mutex
}

func NewCeleryHandler(writer io.Writer, opts HandlerOptions) *CeleryHandler {
	if opts.ProcessName == "" {
		opts.ProcessName = "GoCelery"
	}
	return &CeleryHandler{
		opts:   opts,
		writer: writer,
	}
}

func (h *CeleryHandler) Enabled(ctx context.Context, l Level) bool {
	minLevel := LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}

func (h *CeleryHandler) Handle(ctx context.Context, r Record) error {
	var source string
	if h.opts.AddSource {
		if r.PC != 0 {
			fs := runtime.CallersFrames([]uintptr{r.PC})
			if f, ok := fs.Next(); ok {
				source = fmt.Sprintf(" (%s:%d)", f.File, f.Line)
			}
		}
		if source == "" {
			source = captureSource(5)
		}
	}

	var extra []string
	for _, attr := range h.attrs {
		extra = appendAttr(extra, h.groups, attr)
	}
	r.Attrs(func(attr Attr) bool {
		extra = appendAttr(extra, h.groups, attr)
		return true
	})

	ts := r.Time.Format("2006-01-02 15:04:05,000")
	level := r.Level.String()
	msg := r.Message

	var buf strings.Builder
	buf.WriteString("[")
	buf.WriteString(ts)
	buf.WriteString(": ")
	buf.WriteString(level)
	buf.WriteString("/")
	buf.WriteString(h.opts.ProcessName)
	buf.WriteString("] ")
	buf.WriteString(msg)
	if source != "" {
		buf.WriteString(source)
	}
	if len(extra) > 0 {
		buf.WriteString(" ")
		buf.WriteString(strings.Join(extra, " "))
	}
	buf.WriteString("\n")

	h.mu.Lock()
	_, err := h.writer.Write([]byte(buf.String()))
	h.mu.Unlock()

	return err
}

func captureSource(skip int) string {
	var pcs [10]uintptr
	n := runtime.Callers(skip, pcs[:])
	if n == 0 {
		return ""
	}
	frames := runtime.CallersFrames(pcs[:n])
	for {
		f, ok := frames.Next()
		if !ok {
			return ""
		}
		if !strings.HasPrefix(f.Function, "runtime.") &&
			!strings.HasPrefix(f.Function, "log/slog.") &&
			!strings.Contains(f.Function, "CeleryHandler") {
			return fmt.Sprintf(" (%s:%d)", f.File, f.Line)
		}
	}
}

func appendAttr(dst []string, groups []string, attr Attr) []string {
	val := attr.Value
	if val.Kind() == slog.KindGroup {
		for _, child := range val.Group() {
			dst = appendAttr(dst, append(groups, attr.Key), child)
		}
		return dst
	}

	key := attr.Key
	if len(groups) > 0 {
		key = strings.Join(groups, ".") + "." + key
	}

	dst = append(dst, formatAttr(key, val))
	return dst
}

func formatAttr(key string, val Value) string {
	switch val.Kind() {
	case slog.KindString:
		s := val.String()
		if strings.ContainsAny(s, " \t\n\"") {
			return fmt.Sprintf(`%s=%q`, key, s)
		}
		return fmt.Sprintf("%s=%s", key, s)
	case slog.KindInt64:
		return fmt.Sprintf("%s=%d", key, val.Int64())
	case slog.KindUint64:
		return fmt.Sprintf("%s=%d", key, val.Uint64())
	case slog.KindFloat64:
		return fmt.Sprintf("%s=%f", key, val.Float64())
	case slog.KindBool:
		return fmt.Sprintf("%s=%t", key, val.Bool())
	case slog.KindDuration:
		return fmt.Sprintf("%s=%s", key, val.Duration())
	case slog.KindTime:
		return fmt.Sprintf("%s=%s", key, val.Time().Format(time.RFC3339))
	default:
		v := val.Any()
		if v == nil {
			return fmt.Sprintf("%s=null", key)
		}
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
			b, err := json.Marshal(v)
			if err == nil {
				return fmt.Sprintf("%s=%s", key, string(b))
			}
		}
		return fmt.Sprintf("%s=%v", key, v)
	}
}

func (h *CeleryHandler) WithAttrs(attrs []Attr) Handler {
	newAttrs := make([]Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)
	return &CeleryHandler{
		opts:   h.opts,
		writer: h.writer,
		attrs:  newAttrs,
		groups: h.groups,
	}
}

func (h *CeleryHandler) WithGroup(name string) Handler {
	newGroups := make([]string, len(h.groups)+1)
	copy(newGroups, h.groups)
	newGroups[len(h.groups)] = name
	return &CeleryHandler{
		opts:   h.opts,
		writer: h.writer,
		attrs:  h.attrs,
		groups: newGroups,
	}
}
