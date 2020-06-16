package main //testy

import (
	"fmt"

	generated "./generatedInGo"
)

// T Struct : Currently holds 2 fields, 'name' runs the function and 'result' is the expected return value.
// I might add another field that is of string type, where upon failure the name of the function tested will be provided.
type T struct {
	name   float64
	result float64
}

// RunTest : Takes in a list of 'T' Structs and runs the tests.
func RunTest(x []T) {
	l := len(x)
	fmt.Println("------\nRunning Test Script")
	PassedTests, failed, numtested := 0, 0, 0
	for i := 0; i < l; i++ {
		holder, expected := x[i].name, x[i].result
		if holder == expected {
			PassedTests++
		} else {
			failed++
			fmt.Printf("--TEST FAILED at index: %d.\n", i)
			fmt.Printf("--Expected: %f.\n", expected)
			fmt.Printf("--Actual:   %f.\n", holder)
		}
		numtested++
	}
	fmt.Printf("Ran %d tests out of %d.\n", numtested, l)
	fmt.Printf("Passed %d tests out of %d.\n", PassedTests, l)
	fmt.Println("------")
}

// DoTesting : Runs all of the current tests.
func DoTesting() {
	RunTest(Tester)
}

func main() {
	DoTesting()
}

// Tester : Currently holds all of the tests.
var Tester = []T{
	{
		name:   generated.Max_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 6,
	},
	{
		name:   generated.Max_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	{
		name:   generated.Max_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Max_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 6,
	},
	T{
		name:   generated.Max_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 6,
		///data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Max_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 6,
		//data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Max_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 6,
		/// data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	T{
		name:   generated.Max_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 6,
		//data : []float64{7,5,5,1,4,5,2,2,3,5,6,2,3,3,3,1}
	},
	T{
		name:   generated.Max_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 6,
		//data : []float64{4,4,6,4,1,1,3,4,4,6,6,5,2,2,4,3}
	},
	T{
		name:   generated.Max_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 6,
	},
	T{
		name:   generated.Max_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 5,
	},
	T{
		name:   generated.Max_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 7,
	},
	T{
		name:   generated.Max_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 5,
	},
	T{
		name:   generated.Max_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4,
	},
	T{
		name:   generated.Max_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4,
	},
	T{
		name:   generated.Max_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 2,
	},
	T{
		name:   generated.Max_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 5,
		//data : []float64{1,7,3,4,4,5,5,4,2,2,6,5,4,6,5,7}
	},
	T{
		name:   generated.Max_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 4,
	},
	T{
		name:   generated.Max_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3,
	},
	T{
		name:   generated.Max_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 5,
	},
	T{
		name:   generated.Max_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 3,
	},
	T{
		name:   generated.Max_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 3,
	},
	T{
		name:   generated.Max_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 5,
	},
	T{
		name:   generated.Max_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 1,
	},
	T{
		name:   generated.Max_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 5,
	},
	T{
		name:   generated.Max_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2,
	},
	T{
		name:   generated.Max_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 3,
	},
	T{
		name:   generated.Max_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5,
	},
	T{
		name:   generated.Max_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 3,
	},
	T{
		name:   generated.Max_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 3,
	},
	T{
		name:   generated.Max_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 3,
	},
	T{
		name:   generated.Max_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 2,
	},
	T{
		name:   generated.Max_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 4,
	},
	T{
		name:   generated.Max_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 3,
	},
	T{
		name:   generated.Max_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 3,
	},
	T{
		name:   generated.Max_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 3,
	},
	T{
		name:   generated.Max_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 3,
	},
	T{
		name:   generated.Max_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5,
	},
	T{
		name:   generated.Max_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3,
	},
	T{
		name:   generated.Max_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 4,
	},
	T{
		name:   generated.Max_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 6,
	},
	T{
		name:   generated.Min_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 5,
	},
	T{
		name:   generated.Min_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 3,
		//[]float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Min_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Min_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 5,
	},
	T{
		name:   generated.Min_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3,
	},
	T{
		name:   generated.Min_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3,
	},
	T{
		name:   generated.Min_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1,
	},
	T{
		name:   generated.Min_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 3,
	},
	T{
		name:   generated.Min_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 4,
	},
	T{
		name:   generated.Min_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 3,
	},
	T{
		name:   generated.Min_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3,
	},
	T{
		name:   generated.Min_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 3,
	},
	T{
		name:   generated.Min_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 2,
		//data : float64{7,6,5,6,5,4,1,4,7,5,4,2,5,4,3,3}
	},
	T{
		name:   generated.Min_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Min_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Min_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 1,
		//data : []float64{1,2,3,2,5,6,7,4,1,3,4,6,1,2,4,4}
	},
	T{
		name:   generated.Min_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 3,
		//data : []float64{1,7,3,4,4,5,5,4,2,2,6,5,4,6,5,7}
	},
	T{
		name:   generated.Min_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1,
		//data : []float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Min_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1,
		//data : []float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Min_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1,
		//data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	T{
		name:   generated.Min_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 1,
		//data : []float64{4,4,6,4,1,1,3,4,4,6,6,5,2,2,4,3}
	},
	T{
		name:   generated.Min_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 1,
		//data : []float64{4,3,5,5,2,1,1,2,3,4,6,6,3,1,2,3}
	},
	T{
		name:   generated.Min_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 2,
	},
	T{
		name:   generated.Min_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 1,
		//data : []float64{4,1,3,1,4,6,1,5,5,2,7,2,3,1,6,1}
	},
	T{
		name:   generated.Min_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Min_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2,
		//data : []float64{6,4,4,4,5,2,2,1,3,3,5,4,4,3,3,3}
	},
	T{
		name:   generated.Min_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 1,
	},
	T{
		name:   generated.Min_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2,
		//data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Min_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 2,
	},
	T{
		name:   generated.Min_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1,
		//  data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	T{
		name:   generated.Min_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 2,
		//data : []float64{7,5,5,1,4,5,2,2,3,5,6,2,3,3,3,1}
	},
	T{
		name:   generated.Min_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 1,
	},
	T{
		name:   generated.Min_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 3,
	},
	T{
		name:   generated.Min_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 2,
	},
	T{
		name:   generated.Min_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 2,
	},
	T{
		name:   generated.Min_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 2,
	},
	T{
		name:   generated.Min_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 2,
		//data : []float64{4,4,6,4,1,1,3,4,4,6,6,5,2,2,4,3}
	},
	T{
		name:   generated.Min_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 2,
		//data : []float64{4,3,5,5,2,1,1,2,3,4,6,6,3,1,2,3}
	},
	T{
		name:   generated.Min_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 1,
		//data : []float64{7,1,5,4,4,3,3,4,6,6,2,3,4,2,3,1}
	},
	T{
		name:   generated.Min_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 2,
		//[]float64{1,3,7,4,3,6,6,5,3,3,2,6,5,5,5,7}
	},
	T{
		name:   generated.Min_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 2,
		//data : []float64{4,1,3,1,4,6,1,5,5,2,7,2,3,1,6,1}
	},
	T{
		name:   generated.Sum_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 11,
		//data : float64{7,6,5,6,5,4,1,4,7,5,4,2,5,4,3,3}
	},
	T{
		name:   generated.Sum_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 23,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Sum_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 16,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Sum_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 11,
		//data : []float64{1,2,3,2,5,6,7,4,1,3,4,6,1,2,4,4}
	},
	T{
		name:   generated.Sum_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 21,
		//data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Sum_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 14,
		//data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Sum_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 31,
		//data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	T{
		name:   generated.Sum_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 14,
		//data : []float64{7,5,5,1,4,5,2,2,3,5,6,2,3,3,3,1}
	},
	T{
		name:   generated.Sum_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 16,
		//data : []float64{4,4,6,4,1,1,3,4,4,6,6,5,2,2,4,3}
	},
	T{
		name:   generated.Sum_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 14,
		//data : []float64{4,3,5,5,2,1,1,2,3,4,6,6,3,1,2,3}
	},
	T{
		name:   generated.Sum_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 12,
		//data : []float64{7,1,5,4,4,3,3,4,6,6,2,3,4,2,3,1}
	},
	T{
		name:   generated.Sum_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 16,
		//data : []float64{4,1,3,1,4,6,1,5,5,2,7,2,3,1,6,1}
	},
	T{
		name:   generated.Sum_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 7,
		//data : float64{7,6,5,6,5,4,1,4,7,5,4,2,5,4,3,3}
	},
	T{
		name:   generated.Sum_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 14,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Sum_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 7,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Sum_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 3,
		//data : []float64{1,2,3,2,5,6,7,4,1,3,4,6,1,2,4,4}
	},
	T{
		name:   generated.Sum_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 12,
		//data : []float64{1,7,3,4,4,5,5,4,2,2,6,5,4,6,5,7}
	},
	T{
		name:   generated.Sum_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 12,
		//  data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Sum_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5,
		//data : float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Sum_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 26,
		//data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	T{
		name:   generated.Sum_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 6,
		//data : []float64{4,4,6,4,1,1,3,4,4,6,6,5,2,2,4,3}
	},
	T{
		name:   generated.Sum_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5,
		//data : []float64{4,3,5,5,2,1,1,2,3,4,6,6,3,1,2,3}
	},
	T{
		name:   generated.Sum_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 10,
		//data : []float64{1,3,7,4,3,6,6,5,3,3,2,6,5,5,5,7}
	},
	T{
		name:   generated.Sum_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 3,
		//data : []float64{4,1,3,1,4,6,1,5,5,2,7,2,3,1,6,1}
	},
	T{
		name:   generated.Sum_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9,
		//data : []float64{3,4,2,2,5,6,6,4,4,3,1,1,4,6,4,4}
	},
	T{
		name:   generated.Sum_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 4,
		//data : []float64{6,4,4,4,5,2,2,1,3,3,5,4,4,3,3,3}
	},
	T{
		name:   generated.Sum_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 6,
		//data : []float64{1,7,3,4,4,5,5,4,2,2,6,5,4,6,5,7}
	},
	T{
		name:   generated.Sum_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9,
		//data : []float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
	},
	T{
		name:   generated.Sum_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 5,
		//data : []float64{1,3,3,3,2,5,5,6,4,4,2,3,3,3,4,4}
	},
	T{
		name:   generated.Sum_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 13,
		//data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	T{
		name:   generated.Sum_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 8,
		//data : []float64{7,5,5,1,4,5,2,2,3,5,6,2,3,3,3,1}
	},
	T{
		name:   generated.Sum_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 4,
		//data : []float64{2,3,6,5,7,6,6,4,5,5,4,3,3,6,6,3}
	},
	T{
		name:   generated.Sum_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 7,
		//data : []float64{1,3,3,5,5,5,5,2,4,4,4,3,3,1,5,5}
	},
	T{
		name:   generated.Sum_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 7,
		//data : []float64{2,7,5,5,6,3,7,4,4,5,6,5,3,3,3,5}
	},
	T{
		name:   generated.Sum_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 7,
		//data : []float64{7,1,3,3,2,5,1,4,4,3,2,3,5,5,5,3}
	},
	T{
		name:   generated.Sum_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 11,
		//data : []float64{3,1,1,4,5,5,5,6,2,2,4,4,3,2,1,1}
	},
	T{
		name:   generated.Sum_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 8,
		//data : []float64{4,4,6,4,1,1,3,4,4,6,6,5,2,2,4,3}
	},
	T{
		name:   generated.Sum_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 10,
		//data : []float64{4,3,5,5,2,1,1,2,3,4,6,6,3,1,2,3}
	},
	T{
		name:   generated.Sum_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 6,
		//data : []float64{7,1,5,4,4,3,3,4,6,6,2,3,4,2,3,1}
	},
	T{
		name:   generated.Sum_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 9,
	},
	T{
		name:   generated.Sum_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 11,
	},
	///// RANGE FUNCTION TESTING STARTS HERE \\\\\
	/// MAX_RANGE TESTING \\\
	T{
		name:   generated.Max_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2,
	},
	/*T{
		name:   generated.Max_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 5,
	},*/
	/*T{
			name:   generated.Max_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
			result: 5,
	    },*/ //This function doesnt work
	T{
		name:   generated.Max_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2,
	},
	/// MIN_RANGE TESTING \\\
	T{
		name:   generated.Min_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1,
	},
	/*T{
			name:   generated.Min_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
			result: 2,
	    },*/ //Not Working
	T{
		name:   generated.Min_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 1,
	},
	T{
		name:   generated.Min_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1,
	},
	/// SUM_RANGE TESTING \\\
	T{
		name:   generated.Sum_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9,
	},
	/*T{
				name:   generated.Sum_range_decreasing_sequence(([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4})),
				result: 9,
	    },*/
	/*T{
		name:   generated.Sum_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 10,
	},*/ //Not Working
	T{
		name:   generated.Sum_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9,
	},
	T{
		name:   generated.Sum_range_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9,
	},
	/*T{
		name:   generated.Sum_range_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 9,
	},*/ //Almost working, off by 1 for this test case
}
