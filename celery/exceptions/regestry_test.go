package exceptions

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	e "go_celery_client/celery/internal/errors"
)

var ErrTest = errors.New("test error")

func TestExceptionRegistry_Registered(t *testing.T) {
	r := NewExceptionRegistry()
	if err := r.RegisterException(ErrTest, BaseException{
		ExceptionType:   "ValueError",
		ExceptionModule: "builtins",
	}); err != nil {
		t.Fatal(err)
	}

	info := r.ExceptionInfo(ErrTest, nil, nil, nil)

	if info.ExceptionType != "ValueError" {
		t.Errorf("exc_type = %q, want ValueError", info.ExceptionType)
	}
	if info.ExceptionModule != "builtins" {
		t.Errorf("exc_module = %q, want builtins", info.ExceptionModule)
	}
	if len(info.ExceptionMessage) != 1 || info.ExceptionMessage[0] != "test error" {
		t.Errorf("exc_message = %v, want [test error]", info.ExceptionMessage)
	}
}

func TestExceptionRegistry_WrappedError(t *testing.T) {
	r := NewExceptionRegistry()
	r.RegisterException(ErrTest, BaseException{
		ExceptionType:   "ValueError",
		ExceptionModule: "builtins",
	})

	info := r.ExceptionInfo(fmt.Errorf("wrap: %w", ErrTest), nil, nil, nil)

	if info.ExceptionType != "ValueError" {
		t.Errorf("exc_type = %q, want ValueError", info.ExceptionType)
	}
}

func TestExceptionRegistry_Unregistered(t *testing.T) {
	r := NewExceptionRegistry()

	info := r.ExceptionInfo(errors.New("something else"), nil, nil, nil)

	if info.ExceptionType != "Exception" {
		t.Errorf("exc_type = %q, want Exception", info.ExceptionType)
	}
	if info.ExceptionModule != "builtins" {
		t.Errorf("exc_module = %q, want builtins", info.ExceptionModule)
	}
}

func TestExceptionRegistry_MessageOverride(t *testing.T) {
	r := NewExceptionRegistry()

	info := r.ExceptionInfo(errors.New("x"), []string{"custom message"}, nil, nil)

	if len(info.ExceptionMessage) != 1 || info.ExceptionMessage[0] != "custom message" {
		t.Errorf("exc_message = %v, want [custom message]", info.ExceptionMessage)
	}
}

func TestExceptionRegistry_Duplicate(t *testing.T) {
	r := NewExceptionRegistry()

	if err := r.RegisterException(ErrTest, BaseException{}); err != nil {
		t.Fatal(err)
	}
	if err := r.RegisterException(ErrTest, BaseException{}); err == nil {
		t.Fatal("expected duplicate registration error")
	}
}

func TestExceptionRegistry_Builtin(t *testing.T) {
	r := NewExceptionRegistry()

	info := r.ExceptionInfo(e.ErrNotRegistered, nil, nil, nil)

	if info.ExceptionType != "NotRegistered" {
		t.Errorf("exc_type = %q, want NotRegistered", info.ExceptionType)
	}
	if info.ExceptionModule != "celery.exceptions" {
		t.Errorf("exc_module = %q, want celery.exceptions", info.ExceptionModule)
	}
}

func TestExceptionInfo_Marshal(t *testing.T) {
	info := NewExceptionInfo("ValueError", []string{"boom"}, "builtins", nil, nil)

	data, err := json.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"exc_type":"ValueError","exc_message":["boom"],"exc_module":"builtins"}`
	if string(data) != want {
		t.Errorf("marshaled = %s, want %s", data, want)
	}
}
