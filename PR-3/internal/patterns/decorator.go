package patterns

import "gitlab.com/msv/pr3-calculator-v12/internal/domain"

type LoggingDecorator struct {
	next   domain.Operation
	logger domain.Logger
}

func NewLoggingDecorator(next domain.Operation, logger domain.Logger) domain.Operation {
	return LoggingDecorator{
		next:   next,
		logger: logger,
	}
}

func (d LoggingDecorator) Name() string {
	return d.next.Name()
}

func (d LoggingDecorator) Symbol() string {
	return d.next.Symbol()
}

func (d LoggingDecorator) Execute(left, right float64) (float64, error) {
	d.logger.Printf("start operation=%s left=%v right=%v", d.next.Name(), left, right)
	result, err := d.next.Execute(left, right)
	if err != nil {
		d.logger.Printf("finish operation=%s error=%v", d.next.Name(), err)
		return 0, err
	}
	d.logger.Printf("finish operation=%s result=%v", d.next.Name(), result)
	return result, nil
}
