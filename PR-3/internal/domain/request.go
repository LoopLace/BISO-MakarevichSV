package domain

type Request struct {
	Operation string
	Left      float64
	Right     float64
	Precision int
	Verbose   bool
}
