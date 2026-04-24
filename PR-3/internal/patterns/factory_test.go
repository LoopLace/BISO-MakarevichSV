package patterns

import "testing"

func TestOperationFactoryCreate(t *testing.T) {
	factory := NewOperationFactory()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"add", "add", "add", false},
		{"alias plus", "plus", "add", false},
		{"sub", "sub", "sub", false},
		{"mul", "multiply", "mul", false},
		{"div", "divide", "div", false},
		{"pow", "pow", "pow", false},
		{"invalid", "sqrt", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation, err := factory.Create(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for input %q", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if operation.Name() != tt.want {
				t.Fatalf("operation.Name() = %q, want %q", operation.Name(), tt.want)
			}
		})
	}
}
