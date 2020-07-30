package generatedInGo

// Tests : Struct for testing
type Tests struct {
	Method func([]float64) float64
	Args   Args
	Result float64
	Name   string
}

// Args : Arguments
type Args struct {
	Data []float64
}

// Case0 : First test cases
var Case0 = []Tests{
	{
		Method: Max_max_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 6, Name: "Max_max_decreasing()",
	},
	{
		Method: Max_max_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 6, Name: "Max_max_decreasing_sequence()",
	},
	{
		Method: Max_min_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 4, Name: "Max_min_decreasing()",
	},
	{
		Method: Max_min_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 4, Name: "Max_min_decreasing_sequence()",
	},
	{
		Method: Max_width_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 5, Name: "Max_width_decreasing_sequence()",
	},
	{
		Method: Max_surface_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 10, Name: "Max_surface_decreasing()",
	},
	{
		Method: Max_surface_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 18, Name: "Max_surface_decreasing_sequence()",
	},
	{
		Method: Min_max_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 3, Name: "Min_max_decreasing()",
	},
	{
		Method: Min_max_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 4, Name: "Min_max_decreasing_sequence()",
	},
	{
		Method: Min_min_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 1, Name: "Min_min_decreasing()",
	},
	{
		Method: Min_min_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 1, Name: "Min_min_decreasing_sequence([]float64)",
	},
	{
		Method: Min_width_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 2, Name: "Min_width_decreasing_sequence()",
	},
	{
		Method: Min_surface_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 4, Name: "Min_surface_decreasing()",
	},
	{
		Method: Min_surface_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 6, Name: "Min_surface_decreasing_sequence()",
	},
	{
		Method: Sum_max_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 23, Name: "Sum_max_decreasing([]float64)",
	},
	{
		Method: Sum_max_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 16, Name: "Sum_max_decreasing_sequence([]float64)",
	},
	{
		Method: Sum_min_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 14, Name: "Sum_min_decreasing([]float64)",
	},
	{
		Method: Sum_min_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 7, Name: "Sum_min_decreasing_sequence([]float64)",
	},
	{
		Method: Sum_width_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 9, Name: "Sum_width_decreasing_sequence([]float64)",
	},
	{ //Added test cases
		Method: Sum_surface_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 37, Name: "Sum_surface_decreasing()",
	},
	{
		Method: Sum_surface_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 34, Name: "Sum_surface_decreasing_sequence()",
	},
	{
		Method: Max_range_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 2, Name: "Max_range_decreasing()",
	},
	{
		Method: Max_range_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 5, Name: "Max_range_decreasing_sequence()",
	},
	{
		Method: Min_range_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 1, Name: "Min_range_decreasing()",
	},
	{
		Method: Min_range_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 2, Name: "Min_range_decreasing_sequence([]float64)",
	},
	{
		Method: Sum_range_decreasing, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 9, Name: "Sum_range_decreasing()",
	},
	{
		Method: Sum_range_decreasing_sequence, Args: Args{Data: []float64{3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4}},
		Result: 9, Name: "Sum_range_decreasing_sequence([]float64)",
	},
}

//Case1 : Second tests
var Case1 = []Tests{
	{
		Method: Max_max_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 6, Name: "Max_max_dip_on_increasing_sequence()",
	},
	{
		Method: Max_min_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 2, Name: "Max_min_dip_on_increasing_sequence()",
	},
	{
		Method: Max_surface_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 10, Name: "Max_surface_dip_on_increasing_sequence",
	},
	{
		Method: Min_max_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 5, Name: "Min_max_dip_on_increasing_sequence()",
	},
	{
		Method: Min_min_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 1, Name: "Min_min_dip_on_increasing_sequence()",
	},
	{
		Method: Min_surface_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 9, Name: "Min_surface_dip_on_increasing_sequence",
	},
	{
		Method: Sum_max_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 11, Name: "Sum_max_dip_on_increasing_sequence()",
	},
	{
		Method: Sum_min_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 3, Name: "Sum_min_dip_on_increasing_sequence()",
	},
	{
		Method: Sum_surface_dip_on_increasing_sequence, Args: Args{Data: []float64{1, 2, 3, 2, 5, 6, 7, 4, 1, 3, 4, 6, 1, 2, 4, 4}},
		Result: 19, Name: "Sum_surface_dip_on_increasing_sequence",
	},
}

// Case2 : Third case
var Case2 = []Tests{
	{
		Method: Max_max_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 6, Name: "Max_max_increasing()",
	},
	{
		Method: Max_max_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 6, Name: "Max_max_increasing_sequence()",
	},
	{
		Method: Max_min_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 4, Name: "Max_min_increasing()",
	},
	{
		Method: Max_min_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 3, Name: "Max_min_increasing_sequence()",
	},
	{
		Method: Max_width_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 5, Name: "Max_width_increasing_sequence()",
	},
	{
		Method: Max_surface_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 10, Name: "Max_surface_increasing()",
	},
	{
		Method: Max_surface_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 17, Name: "Max_surface_increasing_sequence()",
	},
	{
		Method: Min_max_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 3, Name: "Min_max_increasing()",
	},
	{
		Method: Min_max_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 3, Name: "Min_max_increasing_sequence()",
	},
	{
		Method: Min_min_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 1, Name: "Min_min_increasing([]float64)",
	},
	{
		Method: Min_min_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 1, Name: "Min_min_increasing_sequence([]float64)",
	},
	{
		Method: Min_width_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 2, Name: "Min_width_increasing_sequence()",
	},
	{
		Method: Min_surface_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 4, Name: "Min_surface_increasing()",
	},
	{
		Method: Min_surface_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 4, Name: "Min_surface_increasing_sequence()",
	},
	{
		Method: Sum_max_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 21, Name: "Sum_max_increasing([]float64)",
	},
	{
		Method: Sum_max_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 14, Name: "Sum_max_increasing_sequence([]float64)",
	},
	{
		Method: Sum_min_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 12, Name: "Sum_min_increasing([]float64)",
	},
	{
		Method: Sum_min_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 5, Name: "Sum_min_increasing_sequence([]float64)",
	},
	{
		Method: Sum_width_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 9, Name: "Sum_width_increasing_sequence([]float64)",
	},
	{
		Method: Sum_surface_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 33, Name: "Sum_surface_increasing()",
	},
	{
		Method: Sum_surface_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 29, Name: "Sum_surface_increasing_sequence()",
	},
	{
		Method: Max_range_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 2, Name: "Max_range_increasing()",
	},
	{
		Method: Max_range_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 5, Name: "Max_range_increasing_sequence()",
	},
	{
		Method: Min_range_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 1, Name: "Min_range_increasing([]float64)",
	},
	{
		Method: Min_range_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 2, Name: "Min_range_increasing_sequence([]float64)",
	},
	{
		Method: Sum_range_increasing, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 9, Name: "Sum_range_increasing([]float64)",
	},
	{
		Method: Sum_range_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 3, 3, 4, 6, 6, 3, 1, 3, 3}},
		Result: 9, Name: "Sum_range_increasing_sequence()",
	},
}

// Case3 : Fourth case
var Case3 = []Tests{
	{
		Method: Max_max_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 6, Name: "Max_max_inflexion()",
	},
	{
		Method: Max_min_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 5, Name: "Max_min_inflexion()",
	},
	{
		Method: Max_width_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 3, Name: "Max_width_inflexion()",
	},
	{
		Method: Max_surface_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 14, Name: "Max_surface_inflexion",
	},
	{
		Method: Min_max_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 1, Name: "Min_max_inflexion()",
	},
	{
		Method: Min_min_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 1, Name: "Min_min_inflexion([]float64)",
	},
	{
		Method: Min_width_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 1, Name: "Min_width_inflexion([]float64)",
	},
	{
		Method: Min_surface_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 1, Name: "Min_surface_inflexion",
	},
	{
		Method: Sum_max_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 31, Name: "Sum_max_inflexion()",
	},
	{
		Method: Sum_min_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 26, Name: "Sum_min_inflexion([]float64)",
	},
	{
		Method: Sum_width_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 13, Name: "Sum_width_inflexion([]float64)",
	},
	{
		Method: Sum_surface_inflexion, Args: Args{Data: []float64{1, 2, 6, 6, 4, 4, 3, 5, 2, 5, 1, 5, 3, 3, 4, 4}},
		Result: 49, Name: "Sum_surface_inflexion",
	},
}

// Case4 : Fifth Case
var Case4 = []Tests{
	{
		Method: Max_max_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 6, Name: "Max_max_peak()",
	},
	{
		Method: Max_width_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 3, Name: "Max_width_peak()",
	},
	{
		Method: Max_surface_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 14, Name: "Max_surface_peak",
	},
	{
		Method: Min_max_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 3, Name: "Min_max_peak()",
	},
	{
		Method: Min_width_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 2, Name: "Min_width_peak([]float64)",
	},
	{
		Method: Min_surface_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 9, Name: "Min_surface_peak",
	},
	{
		Method: Sum_max_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 14, Name: "Sum_max_peak()",
	},
	{
		Method: Sum_width_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 8, Name: "Sum_width_peak([]float64)",
	},
	{
		Method: Sum_surface_peak, Args: Args{Data: []float64{7, 5, 5, 1, 4, 5, 2, 2, 3, 5, 6, 2, 3, 3, 3, 1}},
		Result: 32, Name: "Sum_surface_peak",
	},
}

// Case5 : Sixth Case
var Case5 = []Tests{
	{
		Method: Max_max_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 6, Name: "Max_max_strictly_decreasing_sequence()",
	},
	{
		Method: Max_min_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 3, Name: "Max_min_strictly_decreasing_sequence()",
	},
	{
		Method: Max_width_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 3, Name: "Max_width_strictly_decreasing_sequence()",
	},
	{
		Method: Max_surface_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 13, Name: "Max_surface_strictly_decreasing_sequence",
	},
	{
		Method: Min_max_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 4, Name: "Min_max_strictly_decreasing_sequence()",
	},
	{
		Method: Min_min_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 1, Name: "Min_min_strictly_decreasing_sequence()",
	},
	{
		Method: Min_width_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 2, Name: "Min_width_strictly_decreasing_sequence([]float64)",
	},
	{
		Method: Min_surface_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 7, Name: "Min_surface_strictly_decreasing_sequence",
	},
	{
		Method: Sum_max_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 16, Name: "Sum_max_strictly_decreasing_sequence()",
	},
	{
		Method: Sum_min_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 6, Name: "Sum_min_strictly_decreasing_sequence()",
	},
	{
		Method: Sum_width_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 8, Name: "Sum_width_strictly_decreasing_sequence([]float64)",
	},
	{
		Method: Sum_surface_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 31, Name: "Sum_surface_strictly_decreasing_sequence",
	},
	{
		Method: Max_range_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 5, Name: "Max_range_strictly_decreasing_sequence([]float64)",
	},
	{
		Method: Min_range_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 1, Name: "Min_range_strictly_decreasing_sequence (test case 1)", //Provides correct answer for this test case but does not work properly
	},
	{
		Method: Sum_range_strictly_decreasing_sequence, Args: Args{Data: []float64{4, 4, 6, 4, 1, 1, 3, 4, 4, 6, 6, 5, 2, 2, 4, 3}},
		Result: 10, Name: "Sum_range_strictly_decreasing_sequence([]float64)",
	},
}

// Case6 : Seventh Case
var Case6 = []Tests{
	{
		Method: Max_max_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 6, Name: "Max_max_strictly_increasing_sequence()",
	},
	{
		Method: Max_min_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 3, Name: "Max_min_strictly_increasing_sequence()",
	},
	{
		Method: Max_width_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 5, Name: "Max_width_strictly_increasing_sequence()",
	},
	{
		Method: Max_surface_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 16, Name: "Max_surface_strictly_increasing_sequence",
	},
	{
		Method: Min_max_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 3, Name: "Min_max_strictly_increasing_sequence()",
	},
	{
		Method: Min_min_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 1, Name: "Min_min_strictly_increasing_sequence()",
	},
	{
		Method: Min_width_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 2, Name: "Min_width_strictly_increasing_sequence([]float64)",
	},
	{
		Method: Min_surface_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 6, Name: "Min_surface_strictly_increasing_sequence",
	},
	{
		Method: Sum_max_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 14, Name: "Sum_max_strictly_increasing_sequence()",
	},
	{
		Method: Sum_min_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 5, Name: "Sum_min_strictly_increasing_sequence()",
	},
	{
		Method: Sum_width_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 10, Name: "Sum_width_strictly_increasing_sequence([]float64)",
	},
	{
		Method: Sum_surface_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 30, Name: "Sum_surface_strictly_increasing_sequence",
	},
	{
		Method: Max_range_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 5, Name: "Max_range_strictly_increasing_sequence()",
	},
	{
		Method: Min_range_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 2, Name: "Min_range_strictly_increasing_sequence()",
	},
	{
		Method: Sum_range_strictly_increasing_sequence, Args: Args{Data: []float64{4, 3, 5, 5, 2, 1, 1, 2, 3, 4, 6, 6, 3, 1, 2, 3}},
		Result: 9, Name: "Sum_range_strictly_increasing_sequence([]float64)",
	},
}

// Case7 : Eighth Case
var Case7 = []Tests{
	{
		Method: Max_max_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 5, Name: "Max_max_summit()",
	},
	{
		Method: Max_width_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 3, Name: "Max_width_summit()",
	},
	{
		Method: Max_surface_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 13, Name: "Max_surface_summit",
	},
	{
		Method: Min_max_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 3, Name: "Min_max_summit()",
	},
	{
		Method: Min_width_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 1, Name: "Min_width_summit()",
	},
	{
		Method: Min_surface_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 3, Name: "Min_surface_summit",
	},
	{
		Method: Sum_max_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 12, Name: "Sum_max_summit([]float64)",
	},
	{
		Method: Sum_width_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 6, Name: "Sum_width_summit([]float64)",
	},
	{
		Method: Sum_surface_summit, Args: Args{Data: []float64{7, 1, 5, 4, 4, 3, 3, 4, 6, 6, 2, 3, 4, 2, 3, 1}},
		Result: 23, Name: "Sum_surface_summit",
	},
}

// Case8 : Ninth Case
var Case8 = []Tests{
	{
		Method: Max_max_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 7, Name: "Max_max_zigzag()",
	},
	{
		Method: Max_min_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 1, Name: "Max_min_zigzag()",
	},
	{
		Method: Max_width_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 6, Name: "Max_width_zigzag()",
	},
	{
		Method: Max_surface_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 21, Name: "Max_surface_zigzag",
	},
	{
		Method: Min_max_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 3, Name: "Min_max_zigzag()",
	},
	{
		Method: Min_min_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 1, Name: "Min_min_zigzag([]float64)",
	},
	{
		Method: Min_width_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 2, Name: "Min_width_zigzag()",
	},
	{
		Method: Min_surface_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 5, Name: "Min_surface_zigzag",
	},
	{
		Method: Sum_max_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 16, Name: "Sum_max_zigzag([]float64)",
	},
	{
		Method: Sum_min_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 3, Name: "Sum_min_zigzag([]float64)",
	},
	{
		Method: Sum_width_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 11, Name: "Sum_width_zigzag([]float64)",
	},
	{
		Method: Sum_surface_zigzag, Args: Args{Data: []float64{4, 1, 3, 1, 4, 6, 1, 5, 5, 2, 7, 2, 3, 1, 6, 1}},
		Result: 33, Name: "Sum_surface_zigzag",
	},
}

// Case9 : Tenth Case
var Case9 = []Tests{
	{
		Method: Max_max_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 6, Name: "Max_max_bump_on_decreasing_sequence",
	},
	{
		Method: Max_min_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 5, Name: "Max_min_bump_on_decreasing_sequence()",
	},
	{
		Method: Max_surface_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 16, Name: "Max_surface_bump_on_decreasing_sequence(Data []float64)",
	},
	{
		Method: Min_max_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 5, Name: "Min_max_bump_on_decreasing_sequence()",
	},
	{
		Method: Min_min_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 2, Name: "Min_min_bump_on_decreasing_sequence()",
	},
	{
		Method: Min_surface_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 11, Name: "Min_surface_bump_on_decreasing_sequence(Data []float64)",
	},
	{
		Method: Sum_max_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 11, Name: "Sum_max_bump_on_decreasing_sequence()",
	},
	{
		Method: Sum_min_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 7, Name: "Sum_min_bump_on_decreasing_sequence()",
	},
	{
		Method: Sum_surface_bump_on_decreasing_sequence, Args: Args{Data: []float64{7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3}},
		Result: 27, Name: "Sum_surface_bump_on_decreasing_sequence()",
	},
}

// Case10 : Eleventh Case
var Case10 = []Tests{
	{
		Method: Max_min_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 4, Name: "Max_min_decreasing_terrace",
	},
	{
		Method: Max_width_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 2, Name: "Max_width_decreasing_terrace()",
	},
	{
		Method: Max_surface_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 8, Name: "Max_surface_decreasing_terrace",
	},
	{
		Method: Min_min_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 2, Name: "Min_min_decreasing_terrace",
	},
	{
		Method: Min_width_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 2, Name: "Min_width_decreasing_terrace()",
	},
	{
		Method: Min_surface_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 4, Name: "Min_surface_decreasing_terrace",
	},
	{
		Method: Sum_min_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 6, Name: "Sum_min_decreasing_terrace",
	},
	{
		Method: Sum_width_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 4, Name: "Sum_width_decreasing_terrace([]float64)",
	},
	{
		Method: Sum_surface_decreasing_terrace, Args: Args{Data: []float64{6, 4, 4, 4, 5, 2, 2, 1, 3, 3, 5, 4, 4, 3, 3, 3}},
		Result: 12, Name: "Sum_surface_decreasing_terrace",
	},
}

// Case11 : Twelfth Case
var Case11 = []Tests{
	{
		Method: Max_min_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 5, Name: "Max_min_gorge()",
	},
	{
		Method: Max_width_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 3, Name: "Max_width_gorge()",
	},
	{
		Method: Max_surface_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 11, Name: "Max_surface_gorge(Data []float64)",
	},
	{
		Method: Min_min_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 3, Name: "Min_min_gorge([]float64)",
	},
	{
		Method: Min_width_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 1, Name: "Min_width_gorge([]float64)",
	},
	{
		Method: Min_surface_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 5, Name: "Min_surface_gorge(Data []float64)",
	},
	{
		Method: Sum_min_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 12, Name: "Sum_min_gorge([]float64)",
	},
	{
		Method: Sum_width_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 6, Name: "Sum_width_gorge([]float64)",
	},
	{
		Method: Sum_surface_gorge, Args: Args{Data: []float64{1, 7, 3, 4, 4, 5, 5, 4, 2, 2, 6, 5, 4, 6, 5, 7}},
		Result: 25, Name: "Sum_surface_gorge(Data []float64)",
	},
}

// Case12 : Thirteenth Case
var Case12 = []Tests{
	{
		Method: Max_min_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 5, Name: "Max_min_valley()",
	},
	{
		Method: Max_width_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 4, Name: "Max_width_valley()",
	},
	{
		Method: Max_surface_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 15, Name: "Max_surface_valley",
	},
	{
		Method: Min_min_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 2, Name: "Min_min_valley([]float64)",
	},
	{
		Method: Min_width_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 2, Name: "Min_width_valley()",
	},
	{
		Method: Min_surface_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 7, Name: "Min_surface_valley",
	},
	{
		Method: Sum_min_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 10, Name: "Sum_min_valley([]float64)",
	},
	{
		Method: Sum_width_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 9, Name: "Sum_width_valley([]float64)",
	},
	{
		Method: Sum_surface_valley, Args: Args{Data: []float64{1, 3, 7, 4, 3, 6, 6, 5, 3, 3, 2, 6, 5, 5, 5, 7}},
		Result: 35, Name: "Sum_surface_valley",
	},
}

// Case13 : Fourteenth Case
var Case13 = []Tests{
	{
		Method: Max_min_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 5, Name: "Max_min_increasing_terrace",
	},
	{
		Method: Max_width_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 3, Name: "Max_width_increasing_terrace()",
	},
	{
		Method: Max_surface_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 10, Name: "Max_surface_increasing_terrace",
	},
	{
		Method: Min_min_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 3, Name: "Min_min_increasing_terrace",
	},
	{
		Method: Min_width_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 2, Name: "Min_width_increasing_terrace()",
	},
	{
		Method: Min_surface_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 9, Name: "Min_surface_increasing_terrace",
	},
	{
		Method: Sum_min_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 8, Name: "Sum_min_increasing_terrace",
	},
	{
		Method: Sum_width_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 5, Name: "Sum_width_increasing_terrace([]float64)",
	},
	{
		Method: Sum_surface_increasing_terrace, Args: Args{Data: []float64{1, 3, 3, 3, 2, 5, 5, 6, 4, 4, 2, 3, 3, 3, 4, 4}},
		Result: 19, Name: "Sum_surface_increasing_terrace",
	},
}

// Case14 : Fifthteenth Case
var Case14 = []Tests{
	{
		Method: Max_min_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 5, Name: "Max_min_plain",
	},
	{
		Method: Max_width_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 2, Name: "Max_width_plain()",
	},
	{
		Method: Max_surface_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 6, Name: "Max_surface_plain",
	},
	{
		Method: Min_min_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 3, Name: "Min_min_plain",
	},
	{
		Method: Min_width_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 1, Name: "Min_width_plain([]float64)",
	},
	{
		Method: Min_surface_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 4, Name: "Min_surface_plain",
	},
	{
		Method: Sum_min_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 12, Name: "Sum_min_plain",
	},
	{
		Method: Sum_width_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 4, Name: "Sum_width_plain([]float64)",
	},
	{
		Method: Sum_surface_plain, Args: Args{Data: []float64{2, 3, 6, 5, 7, 6, 6, 4, 5, 5, 4, 3, 3, 6, 6, 3}},
		Result: 15, Name: "Sum_surface_plain",
	},
}

// Case15 : Sixthteenth Case
var Case15 = []Tests{
	{
		Method: Max_width_plateau, Args: Args{Data: []float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}},
		Result: 4, Name: "Max_width_plateau()",
	},
	{
		Method: Min_width_plateau, Args: Args{Data: []float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}},
		Result: 3, Name: "Min_width_plateau([]float64)",
	},
	{
		Method: Sum_width_plateau, Args: Args{Data: []float64{1, 3, 3, 5, 5, 5, 5, 2, 4, 4, 4, 3, 3, 1, 5, 5}},
		Result: 7, Name: "Sum_width_plateau([]float64)",
	},
}

// Case16 : Seventeenth Case
var Case16 = []Tests{
	{
		Method: Max_min_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 5, Name: "Max_min_proper_plain",
	},
	{
		Method: Max_width_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 3, Name: "Max_width_proper_plain()",
	},
	{
		Method: Max_surface_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 10, Name: "Max_surface_proper_plain",
	},
	{
		Method: Min_min_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 3, Name: "Min_min_proper_plain",
	},
	{
		Method: Min_width_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 2, Name: "Min_width_proper_plain([]float64)",
	},
	{
		Method: Min_surface_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 8, Name: "Min_surface_proper_plain",
	},
	{
		Method: Sum_min_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 12, Name: "Sum_min_proper_plain",
	},
	{
		Method: Sum_width_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 7, Name: "Sum_width_proper_plain([]float64)",
	},
	{
		Method: Sum_surface_proper_plain, Args: Args{Data: []float64{2, 7, 5, 5, 6, 3, 7, 4, 4, 5, 6, 5, 3, 3, 3, 5}},
		Result: 27, Name: "Sum_surface_proper_plain",
	},
}

// Case17 : Eighteenth Case
var Case17 = []Tests{
	{
		Method: Max_min_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 5, Name: "Max_min_proper_plateau",
	},
	{
		Method: Max_width_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 3, Name: "Max_width_proper_plateau()",
	},
	{
		Method: Max_surface_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 15, Name: "Max_surface_proper_plateau",
	},
	{
		Method: Min_min_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 3, Name: "Min_min_proper_plateau",
	},
	{
		Method: Min_width_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 2, Name: "Min_width_proper_plateau([]float64)",
	},
	{
		Method: Min_surface_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 6, Name: "Min_surface_proper_plateau",
	},
	{
		Method: Sum_min_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 12, Name: "Sum_min_proper_plateau",
	},
	{
		Method: Sum_width_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 7, Name: "Sum_width_proper_plateau([]float64)",
	},
	{
		Method: Sum_surface_proper_plateau, Args: Args{Data: []float64{7, 1, 3, 3, 2, 5, 1, 4, 4, 3, 2, 3, 5, 5, 5, 3}},
		Result: 29, Name: "Sum_surface_proper_plateau",
	},
}

// Case18 : Nineteenth Case
var Case18 = []Tests{
	{
		Method: Max_min_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 5, Name: "Max_min_steady_sequence", // Added this is referred to as Max_height in catalog
	},
	{
		Method: Max_width_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 3, Name: "Max_width_steady_sequence()",
	},
	{
		Method: Max_surface_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 15, Name: "Max_surface_steady_sequence",
	},
	{
		Method: Min_min_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 1, Name: "Min_min_steady_sequence",
	},
	{
		Method: Min_width_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 2, Name: "Min_width_steady_sequence([]float64)",
	},
	{
		Method: Min_surface_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 2, Name: "Min_surface_steady_sequence",
	},
	{
		Method: Sum_min_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 13, Name: "Sum_min_steady_sequence",
	},
	{
		Method: Sum_width_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 11, Name: "Sum_width_steady_sequence([]float64)",
	},
	{
		Method: Sum_surface_steady_sequence, Args: Args{Data: []float64{3, 1, 1, 4, 5, 5, 5, 6, 2, 2, 4, 4, 3, 2, 1, 1}},
		Result: 31, Name: "Sum_surface_steady_sequence",
	},
}

// Case19 : Twentieth Case
var Case19 = []Tests{
	{
		Method: Max_min_plateau, Args: Args{Data: []float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}},
		Result: 5, Name: "Max_min_plateau",
	},
	{
		Method: Max_surface_plateau, Args: Args{Data: []float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}},
		Result: 10, Name: "Max_surface_plateau",
	},
	{
		Method: Min_min_plateau, Args: Args{Data: []float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}},
		Result: 3, Name: "Min_min_plateau",
	},
	{
		Method: Min_surface_plateau, Args: Args{Data: []float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}},
		Result: 3, Name: "Min_surface_plateau",
	},
	{
		Method: Sum_min_plateau, Args: Args{Data: []float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}},
		Result: 12, Name: "Sum_min_plateau",
	},
	{
		Method: Sum_surface_plateau, Args: Args{Data: []float64{7, 5, 2, 3, 1, 2, 2, 4, 3, 3, 4, 5, 5, 2, 2, 5}},
		Result: 17, Name: "Sum_surface_plateau",
	},
}

// Case20 : Twentyfirst Case
var Case20 = []Tests{
	{
		Method: Max_min_steady, Args: Args{Data: []float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}},
		Result: 6, Name: "Max_min_steady",
	},
	{
		Method: Max_surface_steady, Args: Args{Data: []float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}},
		Result: 12, Name: "Max_surface_steady",
	},
	{
		Method: Min_min_steady, Args: Args{Data: []float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}},
		Result: 1, Name: "Min_min_steady",
	},
	{
		Method: Min_surface_steady, Args: Args{Data: []float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}},
		Result: 2, Name: "Min_surface_steady",
	},
	{
		Method: Sum_min_steady, Args: Args{Data: []float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}},
		Result: 30, Name: "Sum_min_steady",
	},
	{
		Method: Sum_surface_steady, Args: Args{Data: []float64{1, 1, 7, 3, 3, 5, 5, 5, 6, 5, 5, 5, 7, 2, 6, 6}},
		Result: 60, Name: "Sum_surface_steady",
	},
}
