package math

import "testing"

func TestMin(t *testing.T) {
    tests := []struct {
        name     string
        x, y     int
        expected int
    }{
        {"positive numbers", 10, 20, 10},
        {"negative numbers", -10, -20, -20},
        {"equal numbers", 10, 10, 10},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Min(tt.x, tt.y); got != tt.expected {
                t.Errorf("Min() = %v, want %v", got, tt.expected)
            }
        })
    }
}