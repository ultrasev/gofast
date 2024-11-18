package json

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{"simple object", map[string]string{"hello": "world"}},
		{"nested object", map[string]interface{}{
			"name": "test",
			"data": map[string]int{"value": 42},
		}},
		{"array", []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Encode(tt.input)
			if err != nil {
				t.Errorf("Encode() error = %v", err)
				return
			}

			var result interface{}
			if err := Decode(data, &result); err != nil {
				t.Errorf("Decode() error = %v", err)
			}
		})
	}
}

func TestIsValidJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid object", `{"name":"test"}`, true},
		{"valid array", `[1,2,3]`, true},
		{"invalid json", `{"name":}`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidJSON([]byte(tt.input)); got != tt.expected {
				t.Errorf("IsValidJSON() = %v, want %v", got, tt.expected)
			}
		})
	}
}
