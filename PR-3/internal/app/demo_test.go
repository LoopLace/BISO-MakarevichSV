package app

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"gitlab.com/msv/pr3-calculator-v12/internal/patterns"
	"gitlab.com/msv/pr3-calculator-v12/internal/service"
)

func TestRunDemo(t *testing.T) {
	var out bytes.Buffer
	calculator := service.NewCalculator(patterns.NewOperationFactory(), log.New(&bytes.Buffer{}, "", 0))

	err := RunDemo(calculator, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := out.String()
	if !strings.Contains(got, "12.00 + 5.00 = 17.00") {
		t.Fatalf("unexpected demo output: %q", got)
	}
	if !strings.Contains(got, "2 ^ 8 = 256") {
		t.Fatalf("power scenario is missing: %q", got)
	}
}
