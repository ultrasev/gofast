package main

import (
	"fmt"

	"github.com/ultrasev/gofast/pkg/math"
)

func main() {
	result := math.Min(10, 20)
	fmt.Printf("Minimum value: %d\n", result)
}
