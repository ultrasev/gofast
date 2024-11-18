package strings

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"simple string", "hello", "olleh"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "你好世界", "界世好你"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.input); got != tt.expected {
				t.Errorf("Reverse() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple", "hello_world", "helloWorld"},
		{"multiple underscores", "hello_wonderful_world", "helloWonderfulWorld"},
		{"already camel", "helloWorld", "helloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.input); got != tt.expected {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.expected)
			}
		})
	}
}