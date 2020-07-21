package main

import (
	"fmt"
	"math"

	tcp "./TCPServer"
	//generated "./generatedInGo"
)

const (
	connHost = "localhost"
	connPort = "7100"
	connType = "tcp"
)

func add(x float64, y float64) float64  { return (x + y) }
func diff(x float64, y float64) float64 { return (math.Abs(math.Abs(x) - math.Abs(y))) }

func main() {
	tcp.RunServer()
	//fmt.Println(Max_range_decreasing([]float64{14, 13, 11, 9, 5, 4, 2, 1, 0, -1, -3, -5, -8, 25, 24, 20, 18, 16, 14, 13, 11}))
	//fmt.Println(Max_range_increasing_sequence([]float64{24, 26, 28, 29, 30, 34, 37, 20, 24, 26, 28, 29, 30, 34, 37, 20, 24, 26, 28, 29, 30, 34, 37, 20, 24, 26, 28, 29, 30, 34, 37}))
	fmt.Println("=====Simulation has Concluded=====")
}
