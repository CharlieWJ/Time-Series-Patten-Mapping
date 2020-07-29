package main

import (
	"fmt"
	"math"

	tcp "./TCPServer"
	generated "./generatedInGo"
)

const (
	connHost = "localhost"
	connPort = "7100"
	connType = "tcp"
)

func add(x float64, y float64) float64  { return (x + y) }
func diff(x float64, y float64) float64 { return (math.Abs(x - y)) }

func main() {
	tcp.RunServer()
	//fmt.Println(max_range_decreasing([]float64{14, 13, 11, 9, 5, 4, 2, 1, 0, -1, -3, -5, -8, 25, 24, 20, 18, 16, 14, 13, 11}))
	fmt.Println(generated.Sum_width_decreasing_sequence([]float64{4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4, 3}))
	fmt.Println("=====Simulation has Concluded=====")
}

func max_range_decreasing_sequence(data []float64) float64 {
	C := 0.0 //min_f
	D := 0.0 //neutral_f
	R := 0.0 //min_f
	H := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = 0.0
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				H = data[i-1]
				H = math.Max(H, Htemp) // Holding onto the largest value for sequence
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

func max_range_decreasing(data []float64) float64 {
	C := 0.0 //min_f
	D := 0.0 //neutral_f
	R := 0.0 //min_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					D = 0.0                                       //neutral_f
					R = math.Max(Rtemp, diff(data[i-1], data[i])) // R, found_e a0, Range Update
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}
