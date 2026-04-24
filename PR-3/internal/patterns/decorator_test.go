package patterns

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

type mockOperation struct {
	name   string
	symbol string
	result float64
	err    error
}

func (m mockOperation) Name() string   { return m.name }
func (m mockOperation) Symbol() string { return m.symbol }
func (m mockOperation) Execute(_, _ float64) (float64, error) {
	return m.result, m.err
}

func TestLoggingDecoratorExecuteSuccess(t *testing.T) {
	var buffer bytes.Buffer
	logger := log.New(&buffer, "", 0)
	decorated := NewLoggingDecorator(mockOperation{
		name:   "add",
		symbol: "+",
		result: 42,
	}, logger)

	got, err := decorated.Execute(20, 22)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Fatalf("result = %v, want 42", got)
	}

	logOutput := buffer.String()
	if !strings.Contains(logOutput, "start operation=add") {
		t.Fatalf("log output does not contain start message: %q", logOutput)
	}
	if !strings.Contains(logOutput, "finish operation=add result=42") {
		t.Fatalf("log output does not contain finish message: %q", logOutput)
	}
}

func TestLoggingDecoratorExecuteError(t *testing.T) {
	var buffer bytes.Buffer
	logger := log.New(&buffer, "", 0)
	decorated := NewLoggingDecorator(mockOperation{
		name: "div",
		err:  errors.New("boom"),
	}, logger)

	_, err := decorated.Execute(1, 0)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(buffer.String(), "error=boom") {
		t.Fatalf("log output does not contain error details: %q", buffer.String())
	}
}
