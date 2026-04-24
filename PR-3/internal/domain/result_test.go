package domain

import "testing"

func TestResultString(t *testing.T) {
	result := Result{
		Left:      7,
		Right:     2,
		Symbol:    "/",
		Value:     3.5,
		Precision: 2,
	}

	got := result.String()
	want := "7.00 / 2.00 = 3.50"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}
