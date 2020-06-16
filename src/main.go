package main

import (
	"fmt"
	"math"
)

func add(x float64, y float64) float64 {
	return (x + y)
}
func diff(x float64, y float64) float64 { return (math.Abs(x - y)) }

func Max_range_decreasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					D = 0.0
					//fmt.Printf("R: %f  R_temp: %f   D_temp: %f   data[i]: %f  data[i-1]: %f\n", R, R_temp, D_temp, data[i], data[i-1])
					R = math.Max(R_temp, diff(diff(D_temp, data[i-1]), data[i]))
					//fmt.Printf("R: %f\n", R)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				}
			}
			_ = C_temp
			_ = D_temp
			_ = R_temp
		}
	}
	return math.Max(R, C)
}

func Sum_range_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(R_temp, C_temp)
					fmt.Println("R ", R)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = diff(diff(D_temp, data[i-1]), data[i])
					fmt.Println("C ", C)
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(C_temp, diff(D_temp, data[i]))
					fmt.Println("C ", C)
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(R_temp, C_temp)
					fmt.Println("R ", R)
					currentState = 's'
				}
			}
			_ = C_temp
			_ = D_temp
			_ = R_temp
		}
	}
	//fmt.Println("R ", R, "C ", C)
	return add(R, C)
}

func main() {
	fmt.Println(Sum_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}))
	//test.DoTesting()
	//RunTest(tests)
}
