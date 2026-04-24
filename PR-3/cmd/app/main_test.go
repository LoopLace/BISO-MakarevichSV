package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"-op", "add", "-a", "2", "-b", "3"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%q", code, stderr.String())
	}
	if strings.TrimSpace(stdout.String()) != "2.00 + 3.00 = 5.00" {
		t.Fatalf("unexpected stdout: %q", stdout.String())
	}
}

func TestRunHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"-h"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if !strings.Contains(stderr.String(), "Использование:") {
		t.Fatalf("help output not found: %q", stderr.String())
	}
}

func TestRunDemo(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"-demo"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%q", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Демонстрационный сценарий калькулятора") {
		t.Fatalf("demo output not found: %q", stdout.String())
	}
}
