package main

import (
	"fmt"
	"math"
	//generated "./generatedInGo"
)

func add(x float64, y float64) float64 {
	return (x + y)
}
func diff(x float64, y float64) float64 { return (math.Abs(x - y)) }

func Sum_range_decreasing_sequence(data []float64) float64 { //D is essentially a useless variable
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		//fmt.Printf("\n|COUNT %d |\n", i)
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			fmt.Printf("\n|Loop %d |\n", i)
			fmt.Println("----------")
			fmt.Printf("C_temp: %f | D_temp: %f | R_temp: %f |\n\n", C_temp, D_temp, R_temp)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					fmt.Printf("R + C = %f +  %f\n", R_temp, C_temp)
					R = add(R_temp, C_temp)
					//R += D_temp
					fmt.Printf("R = %f\n", R)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = diff(diff(D_temp, data[i-1]), data[i])
					fmt.Println("s: D_temp == ", D_temp)
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = diff(C_temp, diff(D_temp, data[i]))
					fmt.Println("t: D_temp == ", D_temp)
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					//D = diff(D_temp, data[i]) //D is only set here.
					//fmt.Println("Got here")
					currentState = 't'
				}
			}
			_ = C_temp
			_ = D_temp
			_ = R_temp
		}
		//fmt.Println("data[i]: ", data[i])
	}
	return add(R, C)
}

func Max_range_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	temp := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			t_temp := float64(temp)
			fmt.Printf("\n|Loop %d |\n", i)
			fmt.Println("----------")
			fmt.Printf("C_temp: %f | D_temp: %f | R_temp: %f |\n\n", C_temp, D_temp, R_temp)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(R_temp, C_temp)
					//R = math.Max(R, D_temp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					temp = data[i-1]
					C = diff(diff(D_temp, data[i-1]), data[i])
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = diff(C_temp, diff(D_temp, data[i]))
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(D_temp, data[i])
					currentState = 't'
				}
			}
			_ = C_temp
			_ = D_temp
			_ = R_temp
			_ = t_temp
			fmt.Printf("\n\nC_temp: %f | D_temp: %f | R_temp: %f | temp: %f\n\n", C_temp, D_temp, R_temp, temp)
		}
	}
	return math.Max(R, C)
}

func main() {
	//fmt.Println(Sum_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}))
	//																	^2  ^3                  ^1  ^3      ^1
	//fmt.Println(Sum_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4})) //Expected 9
	//test.DoTesting()									     ^2             ^2    ^1 ^2          ^2
	fmt.Println(MaxRangeDecreasingSequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}))
	//fmt.Println(Max_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4})) // Expected 5
}

// MaxRangeDecreasingSequence test
func MaxRangeDecreasingSequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	X := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Xtemp := float64(X)

			fmt.Printf("\n|Loop %d |\n", i)
			fmt.Println("----------")
			fmt.Printf("Ctemp: %f | Dtemp: %f | Rtemp: %f |\n\n", Ctemp, Dtemp, Rtemp)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					//R = math.Max(R, Dtemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					X = data[i-1]
					j := i
					for j < len(data) {
						if X > data[j] {
							Dtemp = D
							D = diff(X, data[j])
							R = math.Max(D, Dtemp)
							fmt.Println(j)
						}
						j++
					}
					C = diff(diff(Dtemp, data[i-1]), data[i])
					//D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i]))
					D = 0.0
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
			_ = Ctemp
			_ = Dtemp
			_ = Rtemp
			_ = Xtemp
			fmt.Printf("\n\nCtemp: %f | Dtemp: %f | Rtemp: %f | Xtemp: %f\n\n", Ctemp, Dtemp, Rtemp, Xtemp)
		}
	}
	return math.Max(R, C)
}
