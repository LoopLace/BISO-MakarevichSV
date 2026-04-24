package patterns

import (
	"errors"
	"testing"
)

func TestStrategiesExecute(t *testing.T) {
	tests := []struct {
		name      string
		operation interface {
			Execute(left, right float64) (float64, error)
		}
		left      float64
		right     float64
		want      float64
		wantError error
	}{
		{"add", AddStrategy{}, 2, 3, 5, nil},
		{"subtract", SubtractStrategy{}, 10, 4, 6, nil},
		{"multiply", MultiplyStrategy{}, 2.5, 4, 10, nil},
		{"divide", DivideStrategy{}, 9, 3, 3, nil},
		{"power", PowerStrategy{}, 2, 3, 8, nil},
		{"division by zero", DivideStrategy{}, 9, 0, 0, ErrDivisionByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.operation.Execute(tt.left, tt.right)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Execute() error = %v, want %v", err, tt.wantError)
			}
			if err == nil && got != tt.want {
				t.Fatalf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
