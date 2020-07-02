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

/*
// Max_range_strictly_increasing_sequence : Exported Function
func Max_range_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	H := math.Inf(1)
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			fmt.Printf("\n|Loop %d |\n", i)
			fmt.Println("----------")
			fmt.Printf("Ctemp: %f | Dtemp: %f | Rtemp: %f |\n\n", Ctemp, Dtemp, Rtemp)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					//C = diff(Ctemp, diff(Dtemp, data[i])) // C, in a0
					C = diff(H, data[i]) // C, in a0  (new)
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp
			_ = Dtemp
			_ = Rtemp
			fmt.Printf("C:     %f | D:     %f | R:     %f \n\n", C, D, R)
		}
	}
	return math.Max(R, C)
}
*/

// Max_range_strictly_increasing_sequence : Exported Function
func Max_range_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	H := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			fmt.Printf("\n| Loop %d |\n", i)
			fmt.Println("----------")
			fmt.Printf("Ctemp: %f | Dtemp: %f | Rtemp: %f | Htemp: %f |\n\n", Ctemp, Dtemp, Rtemp, Htemp)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp
			_ = Dtemp
			_ = Rtemp
			fmt.Printf("C:     %f | D:     %f | R:     %f | H:     %f |\n\n", C, D, R, H)
		}
	}
	return math.Max(R, C)
}

func Max_range_decreasing_sequence(data []float64) float64 { //Do NOT DELETE THIS FUNCTION. WORKING TEMPLATE.
	C := 0.0
	D := 0.0
	R := 0.0
	H := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			Htemp := float64(H)
			//fmt.Printf("\n|Loop %d |\n", i)
			//fmt.Println("----------")
			//fmt.Printf("C_temp: %f | D_temp: %f | R_temp: %f |\n\n", C_temp, D_temp, R_temp)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					H = 0.0
					R = math.Max(R_temp, C_temp)
					//R = math.Max(R, D_temp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				H = data[i-1]          //Gets the largest number and holds it
				H = math.Max(H, Htemp) //The largest number needs to stay the same, and Htemp stored the previous greatest value for comparison
				if currentState == 's' {
					C = diff(diff(D_temp, data[i-1]), data[i])
					//C = diff(diff(0, H), data[i])
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					//C = diff(C_temp, diff(D_temp, data[i]))
					C = diff(H, data[i])
					D = 0.0
					currentState = 't'
					//fmt.Println("Here1")
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(D_temp, data[i])
					currentState = 't'
					//fmt.Println("Here2")
				}
			}
			_ = C_temp
			_ = D_temp
			_ = R_temp
			//fmt.Println("H: ", H)
			//fmt.Printf("\n\nC_temp: %f | D_temp: %f | R_temp: %f \n\n", C_temp, D_temp, R_temp)
			//fmt.Printf("\nC:      %f | D:      %f | R:      %f \n\n", C, D, R)

		}
	}
	return math.Max(R, C)
}

func Max_range_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	H := 0.0
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			Htemp := float64(H)
			//fmt.Printf("\n|Loop %d |\n", i)
			//fmt.Println("----------")
			//fmt.Printf("C_temp: %f | D_temp: %f | R_temp: %f |\n\n", C_temp, D_temp, R_temp)
			if data[i] > data[i-1] {
				H = 0.0
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					//H = 0.0
					R = math.Max(R_temp, C_temp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				H = data[i-1]
				H = math.Max(H, Htemp) //The largest number needs to stay the same, and Htemp stored the previous greatest value for comparison
				if currentState == 's' {
					C = diff(diff(D_temp, data[i-1]), data[i])
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					//C = diff(C_temp, diff(D_temp, data[i]))
					C = diff(H, data[i])
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				H = 0.0 // No longer decreasing
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(R_temp, C_temp)
					currentState = 's'
				}
			}
			_ = C_temp
			_ = D_temp
			_ = R_temp
			//fmt.Printf("C:      %f | D:      %f | R:      %f \n\n", C, D, R)
		}
	}
	return math.Max(R, C)
}

func main() {
	//fmt.Println(Sum_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}))
	//																	^2  ^3                  ^1  ^3      ^1
	//fmt.Println(Sum_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4})) //Expected 9
	//test.DoTesting()									     ^2             ^2    ^1 ^2          ^2
	//fmt.Println(Max_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}))           // Expected 5 NOTE: Working ?
	//fmt.Println(Max_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}))  // Expected 5 Note: Working ?
	//fmt.Println(Max_range_strictly_decreasing_sequence([]float64{10, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3})) // Expected 6 Note: Working ?
	//fmt.Println(Min_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}))           // Expected 2 NOTE: Working ?

	fmt.Println(Max_range_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3})) //, result: 5, name:
}

func Min_range_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	H := 0.0 //math.Inf(1)
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			C_temp := float64(C)
			D_temp := float64(D)
			R_temp := float64(R)
			Htemp := float64(H)
			fmt.Printf("\n|Loop %d |\n", i)
			fmt.Println("----------")
			fmt.Printf("C_temp: %f | D_temp: %f | R_temp: %f |\n\n", C_temp, D_temp, R_temp)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					H = 0.0 //math.Inf(1) //Reset
					R = math.Min(R_temp, C_temp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				H = data[i-1]
				H = math.Max(H, Htemp) //The largest number needs to stay the same, and Htemp stored the previous greatest value for comparison
				if currentState == 's' {
					C = diff(diff(0.0 /*D_temp*/, data[i-1]), data[i])
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					//C = diff(C_temp, diff(D_temp, data[i]))
					C = diff(H, data[i])
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
			fmt.Printf("C:      %f | D:      %f | R:      %f \n\n", C, D, R)
		}
	}
	return math.Min(R, C)
}
