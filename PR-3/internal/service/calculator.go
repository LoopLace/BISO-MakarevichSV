package service

import (
	"gitlab.com/msv/pr3-calculator-v12/internal/domain"
	"gitlab.com/msv/pr3-calculator-v12/internal/patterns"
)

type Calculator struct {
	factory domain.OperationFactory
	logger  domain.Logger
}

func NewCalculator(factory domain.OperationFactory, logger domain.Logger) *Calculator {
	return &Calculator{
		factory: factory,
		logger:  logger,
	}
}

func (c *Calculator) Calculate(request domain.Request) (domain.Result, error) {
	op, err := c.factory.Create(request.Operation)
	if err != nil {
		return domain.Result{}, err
	}

	executable := op
	if request.Verbose && c.logger != nil {
		executable = patterns.NewLoggingDecorator(op, c.logger)
	}

	value, err := executable.Execute(request.Left, request.Right)
	if err != nil {
		return domain.Result{}, err
	}

	precision := request.Precision
	if precision < 0 {
		precision = 0
	}

	return domain.Result{
		Operation: op.Name(),
		Symbol:    op.Symbol(),
		Left:      request.Left,
		Right:     request.Right,
		Value:     value,
		Precision: precision,
	}, nil
}
