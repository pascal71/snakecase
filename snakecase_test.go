package snakecase

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"helloWorld", "hello_world"},
		{"ThisIsATest", "this_is_a_test"},
		{"Already_SnakeCase", "already_snake_case"},
		{"WithNumbers123", "with_numbers123"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual := ToSnakeCase(tt.input)
			if actual != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q; expected %q", tt.input, actual, tt.expected)
			}
		})
	}
}
