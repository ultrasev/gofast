package main

import (
	"fmt"

	"github.com/ultrasev/gofast/pkg/json"
	"github.com/ultrasev/gofast/pkg/math"
)

func main() {
	result := math.Min(10, 20)
	fmt.Printf("Minimum value: %d\n", result)

	response, err := json.Encode(map[string]string{"hello": "world"})
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("JSON response:", string(response))
}
