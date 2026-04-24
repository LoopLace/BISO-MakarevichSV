package service

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"

	"gitlab.com/msv/pr3-calculator-v12/internal/domain"
)

type mockFactory struct {
	operation domain.Operation
	err       error
}

func (m mockFactory) Create(_ string) (domain.Operation, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.operation, nil
}

type mockStrategy struct {
	name     string
	symbol   string
	value    float64
	err      error
	executed bool
}

func (m *mockStrategy) Name() string   { return m.name }
func (m *mockStrategy) Symbol() string { return m.symbol }
func (m *mockStrategy) Execute(_, _ float64) (float64, error) {
	m.executed = true
	if m.err != nil {
		return 0, m.err
	}
	return m.value, nil
}

func TestCalculatorCalculateSuccess(t *testing.T) {
	strategy := &mockStrategy{name: "add", symbol: "+", value: 15}
	calculator := NewCalculator(mockFactory{operation: strategy}, log.New(&bytes.Buffer{}, "", 0))

	result, err := calculator.Calculate(domain.Request{
		Operation: "add",
		Left:      10,
		Right:     5,
		Precision: 3,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strategy.executed {
		t.Fatal("strategy was not executed")
	}
	if result.Value != 15 {
		t.Fatalf("result.Value = %v, want 15", result.Value)
	}
	if result.String() != "10.000 + 5.000 = 15.000" {
		t.Fatalf("unexpected formatted result: %q", result.String())
	}
}

func TestCalculatorCalculateFactoryError(t *testing.T) {
	calculator := NewCalculator(mockFactory{err: errors.New("factory error")}, nil)

	_, err := calculator.Calculate(domain.Request{Operation: "unknown"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCalculatorCalculateVerboseLogging(t *testing.T) {
	strategy := &mockStrategy{name: "mul", symbol: "*", value: 8}
	var buffer bytes.Buffer
	logger := log.New(&buffer, "", 0)

	calculator := NewCalculator(mockFactory{operation: strategy}, logger)
	_, err := calculator.Calculate(domain.Request{
		Operation: "mul",
		Left:      2,
		Right:     4,
		Precision: 0,
		Verbose:   true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buffer.String()
	if !strings.Contains(output, "start operation=mul") {
		t.Fatalf("verbose logging is missing start message: %q", output)
	}
	if !strings.Contains(output, "finish operation=mul result=8") {
		t.Fatalf("verbose logging is missing finish message: %q", output)
	}
}
