package patterns

import (
	"fmt"
	"strings"

	"gitlab.com/msv/pr3-calculator-v12/internal/domain"
)

type operationFactory struct{}

func NewOperationFactory() domain.OperationFactory {
	return operationFactory{}
}

func (operationFactory) Create(kind string) (domain.Operation, error) {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "add", "plus":
		return AddStrategy{}, nil
	case "sub", "minus":
		return SubtractStrategy{}, nil
	case "mul", "multiply":
		return MultiplyStrategy{}, nil
	case "div", "divide":
		return DivideStrategy{}, nil
	case "pow", "power":
		return PowerStrategy{}, nil
	default:
		return nil, fmt.Errorf("неподдерживаемая операция: %s", kind)
	}
}
