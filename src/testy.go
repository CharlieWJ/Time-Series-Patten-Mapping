package main //testy

import (
	"fmt"

	generated "./generatedingo"
	//generated "./generatedInGo"
)

// TestCase Struct : 'method' runs the function, 'result' is the expected return value and name is the function.
type TestCase struct {
	method float64
	result float64
	name   string
}

// RunTest : Takes in a list of 'T' Structs and runs the tests.
func RunTest(x []TestCase) {
	l := len(x)
	fmt.Println("------\nRunning Test Script")
	PassedTests, failed, numtested := 0, 0, 0
	for i := 0; i < l; i++ {
		holder, expected := x[i].method, x[i].result
		tcase := x[i].name
		if holder == expected {
			PassedTests++
		} else {
			failed++
			fmt.Printf("--TEST FAILED: %s.\n", tcase) // Prints the test case's name upon failure.
			//fmt.Printf("--Expected: %f.\n", expected)
			//fmt.Printf("--Actual:   %f.\n", holder)
		}
		numtested++
	}
	fmt.Printf("Ran %d tests out of %d.\n", numtested, l)
	fmt.Printf("Passed %d tests out of %d.\n", PassedTests, l)
	fmt.Println("------")
}

// DoTesting : Runs all of the current tests.
func DoTesting() {
	RunTest(TestCases)
}

func main() {
	DoTesting()
}

// TestCases : holds all current test cases
var TestCases = []TestCase{
	{
		method: generated.Max_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 6, name: "Max_max_bump_on_decreasing_sequence",
	},
	{
		method: generated.Max_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6, name: "Max_max_decreasing()",
	},
	{
		method: generated.Max_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6, name: "Max_max_decreasing_sequence()",
	},
	{
		method: generated.Max_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 6, name: "Max_max_dip_on_increasing_sequence()",
	},
	{
		method: generated.Max_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 6, name: "Max_max_increasing()",
	},
	{
		method: generated.Max_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 6, name: "Max_max_increasing_sequence()",
	},
	{
		method: generated.Max_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 6, name: "Max_max_inflexion()",
		/// data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	{
		method: generated.Max_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 6, name: "Max_max_peak()",
		//data : []float64{7,5,5,1,4,5,2,2,3,5,6,2,3,3,3,1}
	},
	{
		method: generated.Max_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 6, name: "Max_max_strictly_decreasing_sequence()",
	},
	{
		method: generated.Max_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 6, name: "Max_max_strictly_increasing_sequence()",
	},
	{
		method: generated.Max_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 5, name: "Max_max_summit()",
	},
	{
		method: generated.Max_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 7, name: "Max_max_zigzag()",
	},
	{
		method: generated.Max_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 5, name: "Max_min_bump_on_decreasing_sequence()",
	},
	{
		method: generated.Max_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Max_min_decreasing()",
	},
	{
		method: generated.Max_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Max_min_decreasing_sequence()",
	},
	{
		method: generated.Max_min_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 4, name: "Max_min_decreasing_terrace",
	}, // Added this is referred to as Max_height in catalog
	{
		method: generated.Max_min_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 5, name: "Max_min_increasing_terrace",
	}, // Added this is referred to as Max_height in catalog
	{
		method: generated.Max_min_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 5, name: "Max_min_plain",
	}, // Added this is referred to as Max_height in catalog

	{
		method: generated.Max_min_plateau([]float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}),
		result: 5, name: "Max_min_plateau",
	}, // Added this is referred to as Max_height in catalog
	{
		method: generated.Max_min_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 5, name: "Max_min_proper_plain",
	}, // Added this is referred to as Max_height in catalog
	{
		method: generated.Max_min_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 5, name: "Max_min_proper_plateau",
	}, // Added this is referred to as Max_height in catalog
	{
		method: generated.Max_min_steady([]float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}),
		result: 6, name: "Max_min_steady",
	}, // Added this is referred to as Max_height in catalog
	{
		method: generated.Max_min_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 5, name: "Max_min_steady_sequence",
	}, // Added this is referred to as Max_height in catalog

	{
		method: generated.Max_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 2, name: "Max_min_dip_on_increasing_sequence()",
	},
	{
		method: generated.Max_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 5, name: "Max_min_gorge()",
	},
	{
		method: generated.Max_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 4, name: "Max_min_increasing()",
	},
	{
		method: generated.Max_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3, name: "Max_min_increasing_sequence()",
	},
	{
		method: generated.Max_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 5, name: "Max_min_inflexion()",
	},
	{
		method: generated.Max_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 3, name: "Max_min_strictly_decreasing_sequence()",
	},
	{
		method: generated.Max_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 3, name: "Max_min_strictly_increasing_sequence()",
	},
	{
		method: generated.Max_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 5, name: "Max_min_valley()",
	},
	{
		method: generated.Max_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 1, name: "Max_min_zigzag()",
	},
	{
		method: generated.Max_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 5, name: "Max_width_decreasing_sequence()",
	},
	{
		method: generated.Max_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2, name: "Max_width_decreasing_terrace()",
	},
	{
		method: generated.Max_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 3, name: "Max_width_gorge()",
	},
	{
		method: generated.Max_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5, name: "Max_width_increasing_sequence()",
	},
	{
		method: generated.Max_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 3, name: "Max_width_increasing_terrace()",
	},
	{
		method: generated.Max_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 3, name: "Max_width_inflexion()",
	},
	{
		method: generated.Max_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 3, name: "Max_width_peak()",
	},
	{
		method: generated.Max_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 2, name: "Max_width_plain()",
	},
	{
		method: generated.Max_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 4, name: "Max_width_plateau()",
	},
	{
		method: generated.Max_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 3, name: "Max_width_proper_plain()",
	},
	{
		method: generated.Max_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 3, name: "Max_width_proper_plateau()",
	},
	{
		method: generated.Max_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 3, name: "Max_width_steady_sequence()",
	},
	{
		method: generated.Max_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 3, name: "Max_width_strictly_decreasing_sequence()",
	},
	{
		method: generated.Max_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5, name: "Max_width_strictly_increasing_sequence()",
	},
	{
		method: generated.Max_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3, name: "Max_width_summit()",
	},
	{
		method: generated.Max_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 4, name: "Max_width_valley()",
	},
	{
		method: generated.Max_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 6, name: "Max_width_zigzag()",
	},
	//MAX SURFACE TESTING\\
	{
		method: generated.Max_surface_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 16, name: "Max_surface_bump_on_decreasing_sequence(data []float64)",
	},
	{
		method: generated.Max_surface_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 10, name: "Max_surface_decreasing()",
	},
	{
		method: generated.Max_surface_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 18, name: "Max_surface_decreasing_sequence()",
	},
	{
		method: generated.Max_surface_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 8, name: "Max_surface_decreasing_terrace",
	},
	{
		method: generated.Max_surface_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 10, name: "Max_surface_dip_on_increasing_sequence",
	},
	{
		method: generated.Max_surface_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 11, name: "Max_surface_gorge(data []float64)",
	},
	{
		method: generated.Max_surface_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 10, name: "Max_surface_increasing()",
	},
	{
		method: generated.Max_surface_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 17, name: "Max_surface_increasing_sequence()",
	},
	{
		method: generated.Max_surface_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 10, name: "Max_surface_increasing_terrace",
	},
	{
		method: generated.Max_surface_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 14, name: "Max_surface_inflexion",
	},
	{
		method: generated.Max_surface_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 14, name: "Max_surface_peak",
	},
	{
		method: generated.Max_surface_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 6, name: "Max_surface_plain",
	},
	{
		method: generated.Max_surface_plateau([]float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}),
		result: 10, name: "Max_surface_plateau",
	},
	{
		method: generated.Max_surface_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 10, name: "Max_surface_proper_plain",
	},
	{
		method: generated.Max_surface_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 15, name: "Max_surface_proper_plateau",
	},
	{
		method: generated.Max_surface_steady([]float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}),
		result: 12, name: "Max_surface_steady",
	},
	{
		method: generated.Max_surface_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 15, name: "Max_surface_steady_sequence",
	},
	{
		method: generated.Max_surface_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 13, name: "Max_surface_strictly_decreasing_sequence",
	},
	{
		method: generated.Max_surface_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 16, name: "Max_surface_strictly_increasing_sequence",
	},
	{
		method: generated.Max_surface_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 13, name: "Max_surface_summit",
	},
	{
		method: generated.Max_surface_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 15, name: "Max_surface_valley",
	},
	{
		method: generated.Max_surface_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 21, name: "Max_surface_zigzag",
	},
	//MIN FUNCTION TESTING STARTS HERE\\
	{
		method: generated.Min_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 5, name: "Min_max_bump_on_decreasing_sequence()",
	},
	{
		method: generated.Min_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 3, name: "Min_max_decreasing()",
	},
	{
		method: generated.Min_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Min_max_decreasing_sequence()",
	},
	{
		method: generated.Min_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 5, name: "Min_max_dip_on_increasing_sequence()",
	},
	{
		method: generated.Min_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3, name: "Min_max_increasing()",
	},
	{
		method: generated.Min_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3, name: "Min_max_increasing_sequence()",
	},
	{
		method: generated.Min_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_max_inflexion()",
	},
	{
		method: generated.Min_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 3, name: "Min_max_peak()",
	},
	{
		method: generated.Min_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 4, name: "Min_max_strictly_decreasing_sequence()",
	},
	{
		method: generated.Min_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 3, name: "Min_max_strictly_increasing_sequence()",
	},
	{
		method: generated.Min_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3, name: "Min_max_summit()",
	},
	{
		method: generated.Min_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 3, name: "Min_max_zigzag()",
	},
	// MIN HEIGHT TESTING
	{
		method: generated.Min_min_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2, name: "Min_min_decreasing_terrace",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 3, name: "Min_min_increasing_terrace",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 3, name: "Min_min_plain",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_plateau([]float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}),
		result: 3, name: "Min_min_plateau",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 3, name: "Min_min_proper_plain",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 3, name: "Min_min_proper_plateau",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_steady([]float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}),
		result: 1, name: "Min_min_steady",
	}, // Added this is referred to as Min_height in catalog
	{
		method: generated.Min_min_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 1, name: "Min_min_steady_sequence",
	}, // Added this is referred to as Min_height in catalog
	// MIN MIN TESTING
	{
		method: generated.Min_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 2, name: "Min_min_bump_on_decreasing_sequence()",
	},
	{
		method: generated.Min_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1, name: "Min_min_decreasing()",
	},
	{
		method: generated.Min_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1, name: "Min_min_decreasing_sequence([]float64)",
	},
	{
		method: generated.Min_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 1, name: "Min_min_dip_on_increasing_sequence()",
	},
	{
		method: generated.Min_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 3, name: "Min_min_gorge([]float64)",
	},
	{
		method: generated.Min_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1, name: "Min_min_increasing([]float64)",
	},
	{
		method: generated.Min_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1, name: "Min_min_increasing_sequence([]float64)",
	},
	{
		method: generated.Min_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_min_inflexion([]float64)",
	},
	{
		method: generated.Min_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 1, name: "Min_min_strictly_decreasing_sequence()",
	},
	{
		method: generated.Min_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 1, name: "Min_min_strictly_increasing_sequence()",
	},
	{
		method: generated.Min_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 2, name: "Min_min_valley([]float64)",
	},
	{
		method: generated.Min_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 1, name: "Min_min_zigzag([]float64)",
	},
	{
		method: generated.Min_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2, name: "Min_width_decreasing_sequence()",
	},
	{
		method: generated.Min_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2, name: "Min_width_decreasing_terrace()",
	},
	{
		method: generated.Min_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 1, name: "Min_width_gorge([]float64)",
	},
	{
		method: generated.Min_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2, name: "Min_width_increasing_sequence()",
	},
	{
		method: generated.Min_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 2, name: "Min_width_increasing_terrace()",
	},
	{
		method: generated.Min_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_width_inflexion([]float64)",
	},
	{
		method: generated.Min_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 2, name: "Min_width_peak([]float64)",
	},
	{
		method: generated.Min_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 1, name: "Min_width_plain([]float64)",
	},
	{
		method: generated.Min_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 3, name: "Min_width_plateau([]float64)",
	},
	{
		method: generated.Min_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 2, name: "Min_width_proper_plain([]float64)",
	},
	{
		method: generated.Min_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 2, name: "Min_width_proper_plateau([]float64)",
	},
	{
		method: generated.Min_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 2, name: "Min_width_steady_sequence([]float64)",
	},
	{
		method: generated.Min_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 2, name: "Min_width_strictly_decreasing_sequence([]float64)",
	},
	{
		method: generated.Min_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 2, name: "Min_width_strictly_increasing_sequence([]float64)",
	},
	{
		method: generated.Min_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 1, name: "Min_width_summit()",
	},
	{
		method: generated.Min_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 2, name: "Min_width_valley()",
	},
	{
		method: generated.Min_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 2, name: "Min_width_zigzag()",
	},
	//MIN SURFACE TESTING\\
	{
		method: generated.Min_surface_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 11, name: "Min_surface_bump_on_decreasing_sequence(data []float64)",
	},
	{
		method: generated.Min_surface_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Min_surface_decreasing()",
	},
	{
		method: generated.Min_surface_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6, name: "Min_surface_decreasing_sequence()",
	},
	{
		method: generated.Min_surface_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 4, name: "Min_surface_decreasing_terrace",
	},
	{
		method: generated.Min_surface_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 9, name: "Min_surface_dip_on_increasing_sequence",
	},
	{
		method: generated.Min_surface_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 5, name: "Min_surface_gorge(data []float64)",
	},
	{
		method: generated.Min_surface_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 4, name: "Min_surface_increasing()",
	},
	{
		method: generated.Min_surface_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 4, name: "Min_surface_increasing_sequence()",
	},
	{
		method: generated.Min_surface_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 9, name: "Min_surface_increasing_terrace",
	},
	{
		method: generated.Min_surface_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_surface_inflexion",
	},
	{
		method: generated.Min_surface_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 9, name: "Min_surface_peak",
	},
	{
		method: generated.Min_surface_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 4, name: "Min_surface_plain",
	},
	{
		method: generated.Min_surface_plateau([]float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}),
		result: 3, name: "Min_surface_plateau",
	},
	{
		method: generated.Min_surface_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 8, name: "Min_surface_proper_plain",
	},
	{
		method: generated.Min_surface_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 6, name: "Min_surface_proper_plateau",
	},
	{
		method: generated.Min_surface_steady([]float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}),
		result: 2, name: "Min_surface_steady",
	},
	{
		method: generated.Min_surface_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 2, name: "Min_surface_steady_sequence",
	},
	{
		method: generated.Min_surface_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 7, name: "Min_surface_strictly_decreasing_sequence",
	},
	{
		method: generated.Min_surface_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 6, name: "Min_surface_strictly_increasing_sequence",
	},
	{
		method: generated.Min_surface_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3, name: "Min_surface_summit",
	},
	{
		method: generated.Min_surface_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 7, name: "Min_surface_valley",
	},
	{
		method: generated.Min_surface_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 5, name: "Min_surface_zigzag",
	},
	// SUM TESTING
	{
		method: generated.Sum_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 11, name: "Sum_max_bump_on_decreasing_sequence()",
	},
	{
		method: generated.Sum_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 23, name: "Sum_max_decreasing([]float64)",
	},
	{
		method: generated.Sum_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 16, name: "Sum_max_decreasing_sequence([]float64)",
	},
	{
		method: generated.Sum_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 11, name: "Sum_max_dip_on_increasing_sequence()",
	},
	{
		method: generated.Sum_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 21, name: "Sum_max_increasing([]float64)",
	},
	{
		method: generated.Sum_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 14, name: "Sum_max_increasing_sequence([]float64)",
	},
	{
		method: generated.Sum_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 31, name: "Sum_max_inflexion()",
	},
	{
		method: generated.Sum_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 14, name: "Sum_max_peak()",
	},
	{
		method: generated.Sum_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 16, name: "Sum_max_strictly_decreasing_sequence()",
	},
	{
		method: generated.Sum_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 14, name: "Sum_max_strictly_increasing_sequence()",
	},
	{
		method: generated.Sum_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 12, name: "Sum_max_summit([]float64)",
	},
	{
		method: generated.Sum_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 16, name: "Sum_max_zigzag([]float64)",
	},
	// SUM_HEIGHT Tests
	{
		method: generated.Sum_min_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 6, name: "Sum_min_decreasing_terrace",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 8, name: "Sum_min_increasing_terrace",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 12, name: "Sum_min_plain",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_plateau([]float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}),
		result: 12, name: "Sum_min_plateau",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 12, name: "Sum_min_proper_plain",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 12, name: "Sum_min_proper_plateau",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_steady([]float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}),
		result: 30, name: "Sum_min_steady",
	}, // Added this is referred to as sum_height in catalog
	{
		method: generated.Sum_min_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 13, name: "Sum_min_steady_sequence",
	}, // Added this is referred to as Sum_height in catalog
	//SUM MIN TESTS
	{
		method: generated.Sum_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 7, name: "Sum_min_bump_on_decreasing_sequence()",
	},
	{
		method: generated.Sum_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 14, name: "Sum_min_decreasing([]float64)",
	},
	{
		method: generated.Sum_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 7, name: "Sum_min_decreasing_sequence([]float64)",
	},
	{
		method: generated.Sum_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 3, name: "Sum_min_dip_on_increasing_sequence()",
	},
	{
		method: generated.Sum_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 12, name: "Sum_min_gorge([]float64)",
	},
	{
		method: generated.Sum_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 12, name: "Sum_min_increasing([]float64)",
	},
	{
		method: generated.Sum_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5, name: "Sum_min_increasing_sequence([]float64)",
	},
	{
		method: generated.Sum_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 26, name: "Sum_min_inflexion([]float64)",
	},
	{
		method: generated.Sum_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 6, name: "Sum_min_strictly_decreasing_sequence()",
	},
	{
		method: generated.Sum_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5, name: "Sum_min_strictly_increasing_sequence()",
	},
	{
		method: generated.Sum_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 10, name: "Sum_min_valley([]float64)",
	},
	{
		method: generated.Sum_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 3, name: "Sum_min_zigzag([]float64)",
	},
	{
		method: generated.Sum_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9, name: "Sum_width_decreasing_sequence([]float64)",
	},
	{
		method: generated.Sum_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 4, name: "Sum_width_decreasing_terrace([]float64)",
	},
	{
		method: generated.Sum_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 6, name: "Sum_width_gorge([]float64)",
	},
	{
		method: generated.Sum_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9, name: "Sum_width_increasing_sequence([]float64)",
	},
	{
		method: generated.Sum_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 5, name: "Sum_width_increasing_terrace([]float64)",
	},
	{
		method: generated.Sum_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 13, name: "Sum_width_inflexion([]float64)",
	},
	{
		method: generated.Sum_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 8, name: "Sum_width_peak([]float64)",
	},
	{
		method: generated.Sum_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 4, name: "Sum_width_plain([]float64)",
	},
	{
		method: generated.Sum_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 7, name: "Sum_width_plateau([]float64)",
	},
	{
		method: generated.Sum_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 7, name: "Sum_width_proper_plain([]float64)",
	},
	{
		method: generated.Sum_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 7, name: "Sum_width_proper_plateau([]float64)",
	},
	{
		method: generated.Sum_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 11, name: "Sum_width_steady_sequence([]float64)",
	},
	{
		method: generated.Sum_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 8, name: "Sum_width_strictly_decreasing_sequence([]float64)",
	},
	{
		method: generated.Sum_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 10, name: "Sum_width_strictly_increasing_sequence([]float64)",
	},
	{
		method: generated.Sum_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 6, name: "Sum_width_summit([]float64)",
	},
	{
		method: generated.Sum_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 9, name: "Sum_width_valley([]float64)",
	},
	{
		method: generated.Sum_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 11, name: "Sum_width_zigzag([]float64)",
	},
	//SUM SURFACE TESTING\\
	{ //Added test cases
		method: generated.Sum_surface_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 27, name: "Sum_surface_bump_on_decreasing_sequence()",
	},
	{ //Added test cases
		method: generated.Sum_surface_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 37, name: "Sum_surface_decreasing()",
	},
	{
		method: generated.Sum_surface_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 34, name: "Sum_surface_decreasing_sequence()",
	},
	{
		method: generated.Sum_surface_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 12, name: "Sum_surface_decreasing_terrace",
	},
	{
		method: generated.Sum_surface_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 19, name: "Sum_surface_dip_on_increasing_sequence",
	},
	{
		method: generated.Sum_surface_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 25, name: "Sum_surface_gorge(data []float64)",
	},
	{
		method: generated.Sum_surface_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 33, name: "Sum_surface_increasing()",
	},
	{
		method: generated.Sum_surface_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 29, name: "Sum_surface_increasing_sequence()",
	},
	{
		method: generated.Sum_surface_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 19, name: "Sum_surface_increasing_terrace",
	},
	{
		method: generated.Sum_surface_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 49, name: "Sum_surface_inflexion",
	},
	{
		method: generated.Sum_surface_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 32, name: "Sum_surface_peak",
	},
	{
		method: generated.Sum_surface_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 15, name: "Sum_surface_plain",
	},
	{
		method: generated.Sum_surface_plateau([]float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}),
		result: 17, name: "Sum_surface_plateau",
	},
	{
		method: generated.Sum_surface_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 27, name: "Sum_surface_proper_plain",
	},
	{
		method: generated.Sum_surface_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 29, name: "Sum_surface_proper_plateau",
	},
	{
		method: generated.Sum_surface_steady([]float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}),
		result: 60, name: "Sum_surface_steady",
	},
	{
		method: generated.Sum_surface_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 31, name: "Sum_surface_steady_sequence",
	},
	{
		method: generated.Sum_surface_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 31, name: "Sum_surface_strictly_decreasing_sequence",
	},
	{
		method: generated.Sum_surface_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 30, name: "Sum_surface_strictly_increasing_sequence",
	},
	{
		method: generated.Sum_surface_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 23, name: "Sum_surface_summit",
	},
	{
		method: generated.Sum_surface_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 35, name: "Sum_surface_valley",
	},
	{
		method: generated.Sum_surface_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 33, name: "Sum_surface_zigzag",
	},
	///// RANGE FUNCTION TESTING STARTS HERE \\\\\
	/// MAX_RANGE TESTING \\\
	{
		method: generated.Max_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2, name: "Max_range_decreasing()",
	},
	{
		method: generated.Max_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 5, name: "Max_range_decreasing_sequence()",
	}, //This function doesnt work UPDATE: I think it is working
	{
		method: generated.Max_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 5, name: "Max_range_strictly_decreasing_sequence([]float64)",
	},
	{
		method: generated.Max_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2, name: "Max_range_increasing()",
	},
	{
		method: generated.Max_range_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5, name: "Max_range_increasing_sequence()",
	},
	{
		method: generated.Max_range_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5, name: "Max_range_strictly_increasing_sequence()",
	},
	/// MIN_RANGE TESTING \\\
	{
		method: generated.Min_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1, name: "Min_range_decreasing()",
	},
	{
		method: generated.Min_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2, name: "Min_range_decreasing_sequence([]float64)",
	}, //Not Working
	{
		method: generated.Min_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 1, name: "Min_range_strictly_decreasing_sequence (test case 1)",
		//Provides correct answer for this test case but does not work properly
	},
	{
		method: generated.Min_range_strictly_decreasing_sequence([]float64{6, 5, 2, 2, 4, 2}),
		result: 2, name: "Min_range_strictly_decreasing_sequence (test case 2)",
		//Not 100% sure if this is working correctly or not
	},
	{
		method: generated.Min_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1, name: "Min_range_increasing([]float64)",
	},
	{
		method: generated.Min_range_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2, name: "Min_range_increasing_sequence([]float64)",
	},
	{
		method: generated.Min_range_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 2, name: "Min_range_strictly_increasing_sequence()",
	},
	/// SUM_RANGE TESTING \\\
	{
		method: generated.Sum_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9, name: "Sum_range_decreasing()",
	},
	{
		method: generated.Sum_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9, name: "Sum_range_decreasing_sequence([]float64)",
	}, //Not working UPDATE: I think its working
	{
		method: generated.Sum_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 10, name: "Sum_range_strictly_decreasing_sequence([]float64)",
	}, //Not Working UPDATE: I think its working
	{
		method: generated.Sum_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9, name: "Sum_range_increasing([]float64)",
	},
	{
		method: generated.Sum_range_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9, name: "Sum_range_increasing_sequence()",
	},
	{
		method: generated.Sum_range_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 9, name: "Sum_range_strictly_increasing_sequence([]float64)",
	}, //I think is working correctly
}
