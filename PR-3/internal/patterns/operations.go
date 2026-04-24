package patterns

import (
	"errors"
	"math"

	"gitlab.com/msv/pr3-calculator-v12/internal/domain"
)

var ErrDivisionByZero = errors.New("деление на ноль")

type AddStrategy struct{}

func (AddStrategy) Name() string   { return "add" }
func (AddStrategy) Symbol() string { return "+" }
func (AddStrategy) Execute(left, right float64) (float64, error) {
	return left + right, nil
}

type SubtractStrategy struct{}

func (SubtractStrategy) Name() string   { return "sub" }
func (SubtractStrategy) Symbol() string { return "-" }
func (SubtractStrategy) Execute(left, right float64) (float64, error) {
	return left - right, nil
}

type MultiplyStrategy struct{}

func (MultiplyStrategy) Name() string   { return "mul" }
func (MultiplyStrategy) Symbol() string { return "*" }
func (MultiplyStrategy) Execute(left, right float64) (float64, error) {
	return left * right, nil
}

type DivideStrategy struct{}

func (DivideStrategy) Name() string   { return "div" }
func (DivideStrategy) Symbol() string { return "/" }
func (DivideStrategy) Execute(left, right float64) (float64, error) {
	if right == 0 {
		return 0, ErrDivisionByZero
	}
	return left / right, nil
}

type PowerStrategy struct{}

func (PowerStrategy) Name() string   { return "pow" }
func (PowerStrategy) Symbol() string { return "^" }
func (PowerStrategy) Execute(left, right float64) (float64, error) {
	return math.Pow(left, right), nil
}

var _ domain.Operation = AddStrategy{}
var _ domain.Operation = SubtractStrategy{}
var _ domain.Operation = MultiplyStrategy{}
var _ domain.Operation = DivideStrategy{}
var _ domain.Operation = PowerStrategy{}
