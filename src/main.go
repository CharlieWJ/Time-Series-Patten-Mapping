package main

import (
	"fmt"

	generated "./generatedInGo"
)

func main() {
	data := []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}
	fmt.Println(data)
	fmt.Println(generated.Min_min_increasing_sequence(data))
}

func add(x int, y int) int {
	return (x + y)
}
