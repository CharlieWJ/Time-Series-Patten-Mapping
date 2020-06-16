package main

import (
	"fmt"

	generated "./generatedInGo"
)

func add(x float64, y float64) float64 {
	return (x + y)
}

/*
func TestMin() *tuple {
	tup0 := tuple.NewTupleFromItems(1, 2, 3)
	return tup0
}
*/
/*func TestMax() float64 {
	tup0 := tuple.NewTupleFromItems(1, 2, 3)
	return tup0.max()
}*/

func main() {
	data := []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}
	d := []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}
	name := generated.Max_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3})
	fmt.Println(generated.Min_min_increasing_sequence(data))
	fmt.Println(generated.Max_max_decreasing(d))
	fmt.Println(name)
	//test.DoTesting()
	//RunTest(tests)
}
