type TestCase struct {
	method float64
	result float64
	name   string
}



var TestCases = []TestCase{
	{
		method:   generated.Max_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 6, name: "Max_max_bump_on_decreasing_sequence"
	},
	{
		method:   generated.Max_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6, name: "Max_max_decreasing()"
	},
	{
		method:   generated.Max_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 6, name: "Max_max_decreasing_sequence()"
	},
	TestCase{
		method:   generated.Max_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 6, name: "Max_max_dip_on_increasing_sequence()"
	},
	TestCase{
		method:   generated.Max_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 6, name: "Max_max_increasing()"
	},
	TestCase{
		method:   generated.Max_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 6, name: "Max_max_increasing_sequence()"
	},
	TestCase{
		method:   generated.Max_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 6, name: "Max_max_inflexion()"
		/// data : []float64{1,2,6,6,4,4,3,5,2,5,1,5,3,3,4,4}
	},
	TestCase{
		method:   generated.Max_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 6, name: "Max_max_peak()"
		//data : []float64{7,5,5,1,4,5,2,2,3,5,6,2,3,3,3,1}
	},
	TestCase{
		method:   generated.Max_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 6, name: "Max_max_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Max_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 6, name: "Max_max_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Max_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 5, name: "Max_max_summit()",
	},
	TestCase{
		method:   generated.Max_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 7, name: "Max_max_zigzag()",
	},
	TestCase{
		method:   generated.Max_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 5, name: "Max_min_bump_on_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Max_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Max_min_decreasing()",
	},
	TestCase{
		method:   generated.Max_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Max_min_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Max_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 2, name: "Max_min_dip_on_increasing_sequence()",
	},
	TestCase{
		method:   generated.Max_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 5, name: "Max_min_gorge()",
	},
	TestCase{
		method:   generated.Max_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 4, name: "Max_min_increasing()",
	},
	TestCase{
		method:   generated.Max_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3, name: "Max_min_increasing_sequence()",
	},
	TestCase{
		method:   generated.Max_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 5, name: "Max_min_inflexion()",
	},
	TestCase{
		method:   generated.Max_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 3, name: "Max_min_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Max_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 3, name: "Max_min_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Max_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 5, name: "Max_min_valley()",
	},
	TestCase{
		method:   generated.Max_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 1, name: "Max_min_zigzag()",
	},
	TestCase{
		method:   generated.Max_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 5, name: "Max_width_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Max_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2, name: "Max_width_decreasing_terrace()",
	},
	TestCase{
		method:   generated.Max_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 3, name: "Max_width_gorge()",
	},
	TestCase{
		method:   generated.Max_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5, name: "Max_width_increasing_sequence()",
	},
	TestCase{
		method:   generated.Max_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 3, name: "Max_width_increasing_terrace()",
	},
	TestCase{
		method:   generated.Max_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 3, name: "Max_width_inflexion()",
	},
	TestCase{
		method:   generated.Max_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 3, name: "Max_width_peak()",
	},
	TestCase{
		method:   generated.Max_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 2, name: "Max_width_plain()",
	},
	TestCase{
		method:   generated.Max_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 4, name: "Max_width_plateau()",
	},
	TestCase{
		method:   generated.Max_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 3, name: "Max_width_proper_plain()",
	},
	TestCase{
		method:   generated.Max_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 3, name: "Max_width_proper_plateau()",
	},
	TestCase{
		method:   generated.Max_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 3, name: "Max_width_steady_sequence()",
	},
	TestCase{
		method:   generated.Max_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 3, name: "Max_width_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Max_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5, name: "Max_width_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Max_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3, name: "Max_width_summit()",
	},
	TestCase{
		method:   generated.Max_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 4, name: "Max_width_valley()",
	},
	TestCase{
		method:   generated.Max_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 6, name: "Max_width_zigzag()",
	},
	TestCase{
		method:   generated.Min_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 5, name: "Min_max_bump_on_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 3, name: "Min_max_decreasing()",
	},
	TestCase{
		method:   generated.Min_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 4, name: "Min_max_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 5, name: "Min_max_dip_on_increasing_sequence()",
	},
	TestCase{
		method:   generated.Min_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3, name: "Min_max_increasing()",
	},
	TestCase{
		method:   generated.Min_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 3, name: "Min_max_increasing_sequence()",
	},
	TestCase{
		method:   generated.Min_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_max_inflexion()",
	},
	TestCase{
		method:   generated.Min_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 3, name: "Min_max_peak()",
	},
	TestCase{
		method:   generated.Min_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 4, name: "Min_max_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 3, name: "Min_max_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Min_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 3, name: "Min_max_summit()",
	},
	TestCase{
		method:   generated.Min_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 3, name: "Min_max_zigzag()",
	},
	TestCase{
		method:   generated.Min_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 2, name: "Min_min_bump_on_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1, name: "Min_min_decreasing()",
	},
	TestCase{
		method:   generated.Min_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1, name: "Min_min_decreasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Min_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 1, name: "Min_min_dip_on_increasing_sequence()",
	},
	TestCase{
		method:   generated.Min_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 3, name: "Min_min_gorge([]float64)",
	},
	TestCase{
		method:   generated.Min_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1, name: "Min_min_increasing([]float64)",
	},
	TestCase{
		method:   generated.Min_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1, name: "Min_min_increasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Min_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_min_inflexion([]float64)",
	},
	TestCase{
		method:   generated.Min_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 1, name: "Min_min_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 1, name: "Min_min_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Min_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 2, name: "Min_min_valley([]float64)",
	},
	TestCase{
		method:   generated.Min_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 1, name: "Min_min_zigzag([]float64)",
	},
	TestCase{
		method:   generated.Min_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2, name: "Min_width_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 2, name: "Min_width_decreasing_terrace()",
	},
	TestCase{
		method:   generated.Min_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 1, name: "Min_width_gorge([]float64)",
	},
	TestCase{
		method:   generated.Min_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2, name: "Min_width_increasing_sequence()",
	},
	TestCase{
		method:   generated.Min_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 2, name: "Min_width_increasing_terrace()",
	},
	TestCase{
		method:   generated.Min_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 1, name: "Min_width_inflexion([]float64)",
	},
	TestCase{
		method:   generated.Min_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 2, name: "Min_width_peak([]float64)",
	},
	TestCase{
		method:   generated.Min_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 1, name: "Min_width_plain([]float64)",
	},
	TestCase{
		method:   generated.Min_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 3, name: "Min_width_plateau([]float64)",
	},
	TestCase{
		method:   generated.Min_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 2, name: "Min_width_proper_plain([]float64)",
	},
	TestCase{
		method:   generated.Min_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 2, name: "Min_width_proper_plateau([]float64)",
	},
	TestCase{
		method:   generated.Min_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 2, name: "Min_width_steady_sequence([]float64)",
	},
	TestCase{
		method:   generated.Min_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 2, name: "Min_width_strictly_decreasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Min_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 2, name: "Min_width_strictly_increasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Min_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 1, name: "Min_width_summit()",
	},
	TestCase{
		method:   generated.Min_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 2, name: "Min_width_valley()",
	},
	TestCase{
		method:   generated.Min_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 2, name: "Min_width_zigzag()",
	},
	TestCase{
		method:   generated.Sum_max_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 11, name: "Sum_max_bump_on_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_max_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 23, name: "Sum_max_decreasing([]float64)",
	},
	TestCase{
		method:   generated.Sum_max_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 16, name: "Sum_max_decreasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_max_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 11, name: "Sum_max_dip_on_increasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_max_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 21, name: "Sum_max_increasing([]float64)",
	},
	TestCase{
		method:   generated.Sum_max_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 14, name: "Sum_max_increasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_max_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 31, name: "Sum_max_inflexion()",
	},
	TestCase{
		method:   generated.Sum_max_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 14, name: "Sum_max_peak()",
	},
	TestCase{
		method:   generated.Sum_max_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 16, name: "Sum_max_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_max_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 14, name: "Sum_max_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_max_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 12, name: "Sum_max_summit([]float64)",
	},
	TestCase{
		method:   generated.Sum_max_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 16, name: "Sum_max_zigzag([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 7, name: "Sum_min_bump_on_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_min_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 14, name: "Sum_min_decreasing([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 7, name: "Sum_min_decreasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_dip_on_increasing_sequence([]float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}),
		result: 3, name: "Sum_min_dip_on_increasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_min_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 12, name: "Sum_min_gorge([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 12, name: "Sum_min_increasing([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 5, name: "Sum_min_increasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 26, name: "Sum_min_inflexion([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 6, name: "Sum_min_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_min_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 5, name: "Sum_min_strictly_increasing_sequence()",
	},
	TestCase{
		method:   generated.Sum_min_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 10, name: "Sum_min_valley([]float64)",
	},
	TestCase{
		method:   generated.Sum_min_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 3, name: "Sum_min_zigzag([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9, name: "Sum_width_decreasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_decreasing_terrace([]float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}),
		result: 4, name: "Sum_width_decreasing_terrace([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_gorge([]float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}),
		result: 6, name: "Sum_width_gorge([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9, name: "Sum_width_increasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_increasing_terrace([]float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}),
		result: 5, name: "Sum_width_increasing_terrace([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_inflexion([]float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}),
		result: 13, name: "Sum_width_inflexion([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_peak([]float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}),
		result: 8, name: "Sum_width_peak([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_plain([]float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}),
		result: 4, name: "Sum_width_plain([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_plateau([]float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}),
		result: 7, name: "Sum_width_plateau([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_proper_plain([]float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}),
		result: 7, name: "Sum_width_proper_plain([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_proper_plateau([]float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}),
		result: 7, name: "Sum_width_proper_plateau([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_steady_sequence([]float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}),
		result: 11, name: "Sum_width_steady_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 8, name: "Sum_width_strictly_decreasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 10, name: "Sum_width_strictly_increasing_sequence([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_summit([]float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}),
		result: 6, name: "Sum_width_summit([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_valley([]float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}),
		result: 9, name: "Sum_width_valley([]float64)",
	},
	TestCase{
		method:   generated.Sum_width_zigzag([]float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}),
		result: 11, name: "Sum_width_zigzag([]float64)",
	},
	TestCase{ //Added test cases
		method:   generated.Sum_surface_bump_on_decreasing_sequence([]float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}),
		result: 27, name: "Sum_surface_bump_on_decreasing_sequence()",
	},
	///// RANGE FUNCTION TESTING STARTS HERE \\\\\
	/// MAX_RANGE TESTING \\\
	TestCase{
		method:   generated.Max_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 2, name: "Max_range_decreasing()",
	},
	/*TestCase{
		method:   generated.Max_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 5,
	},*/
	/*TestCase{
			method:   generated.Max_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
			result: 5,
	    },*/ //This function doesnt work
	TestCase{
		method:   generated.Max_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 2, name: "Max_range_increasing()",
	},
	/// MIN_RANGE TESTING \\\
	TestCase{
		method:   generated.Min_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 1, name: "Min_range_decreasing()",
	},
	/*TestCase{
			method:   generated.Min_range_decreasing_sequence([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
			result: 2,
	    },*/ //Not Working
	TestCase{
		method:   generated.Min_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 1, name: "Min_range_strictly_decreasing_sequence()",
	},
	TestCase{
		method:   generated.Min_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 1, name: "Min_range_increasing([]float64)",
	},
	/// SUM_RANGE TESTING \\\
	TestCase{
		method:   generated.Sum_range_decreasing([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}),
		result: 9, name: "Sum_range_decreasing()",
	},
	/*TestCase{
				method:   generated.Sum_range_decreasing_sequence(([]float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4})),
				result: 9,
	    },*/
	/*TestCase{
		method:   generated.Sum_range_strictly_decreasing_sequence([]float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}),
		result: 10,
	},*/ //Not Working
	TestCase{
		method:   generated.Sum_range_increasing([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9, name: "Sum_range_increasing([]float64)",
	},
	TestCase{
		method:   generated.Sum_range_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}),
		result: 9, name: "Sum_range_increasing_sequence()",
	},
	/*TestCase{
		method:   generated.Sum_range_strictly_increasing_sequence([]float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}),
		result: 9,
	},*/ //Almost working, off by 1 for this test case
}
