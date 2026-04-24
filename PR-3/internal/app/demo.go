package app

import (
	"fmt"
	"io"

	"gitlab.com/msv/pr3-calculator-v12/internal/domain"
	"gitlab.com/msv/pr3-calculator-v12/internal/service"
)

func RunDemo(calculator *service.Calculator, out io.Writer) error {
	scenarios := []domain.Request{
		{Operation: "add", Left: 12, Right: 5, Precision: 2},
		{Operation: "sub", Left: 10, Right: 3, Precision: 2},
		{Operation: "mul", Left: 4, Right: 2.5, Precision: 2},
		{Operation: "div", Left: 7, Right: 2, Precision: 3},
		{Operation: "pow", Left: 2, Right: 8, Precision: 0},
	}

	fmt.Fprintln(out, "Демонстрационный сценарий калькулятора:")
	for _, scenario := range scenarios {
		result, err := calculator.Calculate(scenario)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "- %s\n", result.String())
	}
	return nil
}
