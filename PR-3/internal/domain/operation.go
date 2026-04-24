package domain

type Operation interface {
	Name() string
	Symbol() string
	Execute(left, right float64) (float64, error)
}

type OperationFactory interface {
	Create(kind string) (Operation, error)
}
