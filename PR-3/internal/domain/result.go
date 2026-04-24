package domain

import "fmt"

type Result struct {
	Operation string
	Symbol    string
	Left      float64
	Right     float64
	Value     float64
	Precision int
}

func (r Result) String() string {
	return fmt.Sprintf("%.*f %s %.*f = %.*f", r.Precision, r.Left, r.Symbol, r.Precision, r.Right, r.Precision, r.Value)
}
