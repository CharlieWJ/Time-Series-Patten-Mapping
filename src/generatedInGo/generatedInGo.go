// ----------------------------------------------------------------------------
// This file was auto-generated on 2020-07-27
// By Charles W. Jeffries.
// Source Code : https://github.com/CharlieWJ/Time-Series-Patten-Mapping
// ----------------------------------------------------------------------------

package generatedInGo

import (
	"math"
)

func add(x float64, y float64) float64  { return (x + y) }
func diff(x float64, y float64) float64 { return math.Abs(x - y) } // The absolute difference

// Max_one_bump_on_decreasing_sequence : Exported Function
func Max_one_bump_on_decreasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
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

// Max_one_decreasing : Exported Function
func Max_one_decreasing(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
					D = 1.0
					R = math.Max(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
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

// Max_one_decreasing_sequence : Exported Function
func Max_one_decreasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
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

// Max_one_decreasing_terrace : Exported Function
func Max_one_decreasing_terrace(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_dip_on_increasing_sequence : Exported Function
func Max_one_dip_on_increasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
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

// Max_one_gorge : Exported Function
func Max_one_gorge(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_one_increasing : Exported Function
func Max_one_increasing(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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

// Max_one_increasing_sequence : Exported Function
func Max_one_increasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
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

// Max_one_increasing_terrace : Exported Function
func Max_one_increasing_terrace(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_inflexion : Exported Function
func Max_one_inflexion(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_peak : Exported Function
func Max_one_peak(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_plain : Exported Function
func Max_one_plain(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_plateau : Exported Function
func Max_one_plateau(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_proper_plain : Exported Function
func Max_one_proper_plain(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_proper_plateau : Exported Function
func Max_one_proper_plateau(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_steady : Exported Function
func Max_one_steady(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 1.0
					R = math.Max(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
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

// Max_one_steady_sequence : Exported Function
func Max_one_steady_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_one_strictly_decreasing_sequence : Exported Function
func Max_one_strictly_decreasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_one_strictly_increasing_sequence : Exported Function
func Max_one_strictly_increasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_one_summit : Exported Function
func Max_one_summit(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_one_valley : Exported Function
func Max_one_valley(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
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

// Max_one_zigzag : Exported Function
func Max_one_zigzag(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'e'
				} else if currentState == 'e' {
					D = 1.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'c' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'f' {
					C = 1.0
					D = 1.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_width_bump_on_decreasing_sequence : Exported Function
func Max_width_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
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

// Max_width_decreasing : Exported Function
func Max_width_decreasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
					D = 0.0
					R = math.Max(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
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

// Max_width_decreasing_sequence : Exported Function
func Max_width_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_decreasing_terrace : Exported Function
func Max_width_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_dip_on_increasing_sequence : Exported Function
func Max_width_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
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

// Max_width_gorge : Exported Function
func Max_width_gorge(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_width_increasing : Exported Function
func Max_width_increasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Max(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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

// Max_width_increasing_sequence : Exported Function
func Max_width_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_increasing_terrace : Exported Function
func Max_width_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_inflexion : Exported Function
func Max_width_inflexion(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_peak : Exported Function
func Max_width_peak(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_plain : Exported Function
func Max_width_plain(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_plateau : Exported Function
func Max_width_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_proper_plain : Exported Function
func Max_width_proper_plain(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_proper_plateau : Exported Function
func Max_width_proper_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_steady : Exported Function
func Max_width_steady(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Max(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
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

// Max_width_steady_sequence : Exported Function
func Max_width_steady_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_width_strictly_decreasing_sequence : Exported Function
func Max_width_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_width_strictly_increasing_sequence : Exported Function
func Max_width_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_width_summit : Exported Function
func Max_width_summit(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_width_valley : Exported Function
func Max_width_valley(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
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

// Max_width_zigzag : Exported Function
func Max_width_zigzag(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = add(Dtemp, 1.0)
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = add(Dtemp, 1.0)
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_surface_bump_on_decreasing_sequence : Exported Function
func Max_surface_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
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

// Max_surface_decreasing : Exported Function
func Max_surface_decreasing(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
					D = 0.0
					R = math.Max(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_surface_decreasing_sequence : Exported Function
func Max_surface_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i])
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

// Max_surface_decreasing_terrace : Exported Function
func Max_surface_decreasing_terrace(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_dip_on_increasing_sequence : Exported Function
func Max_surface_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
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

// Max_surface_gorge : Exported Function
func Max_surface_gorge(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_surface_increasing : Exported Function
func Max_surface_increasing(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Max(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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

// Max_surface_increasing_sequence : Exported Function
func Max_surface_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i])
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

// Max_surface_increasing_terrace : Exported Function
func Max_surface_increasing_terrace(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_inflexion : Exported Function
func Max_surface_inflexion(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_peak : Exported Function
func Max_surface_peak(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_plain : Exported Function
func Max_surface_plain(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_plateau : Exported Function
func Max_surface_plateau(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_proper_plain : Exported Function
func Max_surface_proper_plain(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_proper_plateau : Exported Function
func Max_surface_proper_plateau(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Max(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_steady : Exported Function
func Max_surface_steady(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Max(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_surface_steady_sequence : Exported Function
func Max_surface_steady_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_surface_strictly_decreasing_sequence : Exported Function
func Max_surface_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_surface_strictly_increasing_sequence : Exported Function
func Max_surface_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_surface_summit : Exported Function
func Max_surface_summit(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_surface_valley : Exported Function
func Max_surface_valley(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
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

// Max_surface_zigzag : Exported Function
func Max_surface_zigzag(data []float64) float64 {
	C := math.Inf(-1)
	D := 0.0
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = add(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = add(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(-1)
					D = 0.0
					R = math.Max(Rtemp, Ctemp)
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

// Max_max_bump_on_decreasing_sequence : Exported Function
func Max_max_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
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

// Max_max_decreasing : Exported Function
func Max_max_decreasing(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_max_decreasing_sequence : Exported Function
func Max_max_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
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

// Max_max_decreasing_terrace : Exported Function
func Max_max_decreasing_terrace(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_dip_on_increasing_sequence : Exported Function
func Max_max_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
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

// Max_max_gorge : Exported Function
func Max_max_gorge(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_max_increasing : Exported Function
func Max_max_increasing(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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

// Max_max_increasing_sequence : Exported Function
func Max_max_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
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

// Max_max_increasing_terrace : Exported Function
func Max_max_increasing_terrace(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_inflexion : Exported Function
func Max_max_inflexion(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_peak : Exported Function
func Max_max_peak(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_plain : Exported Function
func Max_max_plain(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_plateau : Exported Function
func Max_max_plateau(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_proper_plain : Exported Function
func Max_max_proper_plain(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_proper_plateau : Exported Function
func Max_max_proper_plateau(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_steady : Exported Function
func Max_max_steady(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = math.Inf(-1)
					R = math.Max(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_max_steady_sequence : Exported Function
func Max_max_steady_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_max_strictly_decreasing_sequence : Exported Function
func Max_max_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
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

// Max_max_strictly_increasing_sequence : Exported Function
func Max_max_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
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

// Max_max_summit : Exported Function
func Max_max_summit(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_max_valley : Exported Function
func Max_max_valley(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
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

// Max_max_zigzag : Exported Function
func Max_max_zigzag(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(-1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = math.Inf(-1)
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(-1)
					D = math.Inf(-1)
					R = math.Max(Rtemp, Ctemp)
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

// Max_min_bump_on_decreasing_sequence : Exported Function
func Max_min_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
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

// Max_min_decreasing : Exported Function
func Max_min_decreasing(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_min_decreasing_sequence : Exported Function
func Max_min_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i])
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

// Max_min_decreasing_terrace : Exported Function
func Max_min_decreasing_terrace(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_dip_on_increasing_sequence : Exported Function
func Max_min_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
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

// Max_min_gorge : Exported Function
func Max_min_gorge(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_min_increasing : Exported Function
func Max_min_increasing(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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

// Max_min_increasing_sequence : Exported Function
func Max_min_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i])
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

// Max_min_increasing_terrace : Exported Function
func Max_min_increasing_terrace(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_inflexion : Exported Function
func Max_min_inflexion(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_peak : Exported Function
func Max_min_peak(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_plain : Exported Function
func Max_min_plain(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_plateau : Exported Function
func Max_min_plateau(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_proper_plain : Exported Function
func Max_min_proper_plain(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_proper_plateau : Exported Function
func Max_min_proper_plateau(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_steady : Exported Function
func Max_min_steady(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = math.Inf(1)
					R = math.Max(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_min_steady_sequence : Exported Function
func Max_min_steady_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_min_strictly_decreasing_sequence : Exported Function
func Max_min_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
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

// Max_min_strictly_increasing_sequence : Exported Function
func Max_min_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
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

// Max_min_summit : Exported Function
func Max_min_summit(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_min_valley : Exported Function
func Max_min_valley(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
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
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
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

// Max_min_zigzag : Exported Function
func Max_min_zigzag(data []float64) float64 {
	C := math.Inf(-1)
	D := math.Inf(1)
	R := math.Inf(-1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = math.Inf(1)
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(-1)
					D = math.Inf(1)
					R = math.Max(Rtemp, Ctemp)
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

// Max_range_bump_on_decreasing_sequence : Exported Function
func Max_range_bump_on_decreasing_sequence(data []float64) float64 {
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
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

// Max_range_decreasing : Exported Function
func Max_range_decreasing(data []float64) float64 {
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

// Max_range_decreasing_sequence : Exported Function
func Max_range_decreasing_sequence(data []float64) float64 {
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

// Max_range_decreasing_terrace : Exported Function
func Max_range_decreasing_terrace(data []float64) float64 {
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_dip_on_increasing_sequence : Exported Function
func Max_range_dip_on_increasing_sequence(data []float64) float64 {
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
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
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

// Max_range_gorge : Exported Function
func Max_range_gorge(data []float64) float64 {
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
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_range_increasing : Exported Function
func Max_range_increasing(data []float64) float64 {
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
					D = 0.0                                       //neutral_f
					R = math.Max(Rtemp, diff(data[i-1], data[i])) // R, found_e a0, Range Update
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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

// Max_range_increasing_sequence : Exported Function
func Max_range_increasing_sequence(data []float64) float64 {
	C := 0.0 //min_f
	D := 0.0 //neutral_f
	R := 0.0 //min_f
	H := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
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

// Max_range_increasing_terrace : Exported Function
func Max_range_increasing_terrace(data []float64) float64 {
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
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_inflexion : Exported Function
func Max_range_inflexion(data []float64) float64 {
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
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_peak : Exported Function
func Max_range_peak(data []float64) float64 {
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
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_plain : Exported Function
func Max_range_plain(data []float64) float64 {
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
				} else if currentState == 'r' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_plateau : Exported Function
func Max_range_plateau(data []float64) float64 {
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
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_proper_plain : Exported Function
func Max_range_proper_plain(data []float64) float64 {
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_proper_plateau : Exported Function
func Max_range_proper_plateau(data []float64) float64 {
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
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Max(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_steady : Exported Function
func Max_range_steady(data []float64) float64 {
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0                                                    //neutral_f
					R = math.Max(Rtemp, diff(diff(Dtemp, data[i-1]), data[i])) // R, found_e a0
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

// Max_range_steady_sequence : Exported Function
func Max_range_steady_sequence(data []float64) float64 {
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
				} else if currentState == 'r' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(Ctemp, diff(Dtemp, data[i])) // C, in a0
					D = 0.0                               //neutral_f
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_range_strictly_decreasing_sequence : Exported Function
func Max_range_strictly_decreasing_sequence(data []float64) float64 {
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
				} else if currentState == 'r' {
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
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				H = 0.0
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
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

// Max_range_strictly_increasing_sequence : Exported Function
func Max_range_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0 //min_f
	D := 0.0 //neutral_f
	R := 0.0 //min_f
	H := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
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

// Max_range_summit : Exported Function
func Max_range_summit(data []float64) float64 {
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
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Max(R, C)
}

// Max_range_valley : Exported Function
func Max_range_valley(data []float64) float64 {
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
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
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

// Max_range_zigzag : Exported Function
func Max_range_zigzag(data []float64) float64 {
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
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = diff(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0 //neutral_f
					currentState = 'a'
				} else if currentState == 'f' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = diff(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0 //min_f
					D = 0.0 //neutral_f
					R = math.Max(Rtemp, Ctemp)
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

// Min_one_bump_on_decreasing_sequence : Exported Function
func Min_one_bump_on_decreasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_decreasing : Exported Function
func Min_one_decreasing(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
					D = 1.0
					R = math.Min(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
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
	return math.Min(R, C)
}

// Min_one_decreasing_sequence : Exported Function
func Min_one_decreasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_decreasing_terrace : Exported Function
func Min_one_decreasing_terrace(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_dip_on_increasing_sequence : Exported Function
func Min_one_dip_on_increasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_gorge : Exported Function
func Min_one_gorge(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_increasing : Exported Function
func Min_one_increasing(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return math.Min(R, C)
}

// Min_one_increasing_sequence : Exported Function
func Min_one_increasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_increasing_terrace : Exported Function
func Min_one_increasing_terrace(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_inflexion : Exported Function
func Min_one_inflexion(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_peak : Exported Function
func Min_one_peak(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_plain : Exported Function
func Min_one_plain(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_plateau : Exported Function
func Min_one_plateau(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_proper_plain : Exported Function
func Min_one_proper_plain(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_proper_plateau : Exported Function
func Min_one_proper_plateau(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_steady : Exported Function
func Min_one_steady(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 1.0
					R = math.Min(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_steady_sequence : Exported Function
func Min_one_steady_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_strictly_decreasing_sequence : Exported Function
func Min_one_strictly_decreasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_strictly_increasing_sequence : Exported Function
func Min_one_strictly_increasing_sequence(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_summit : Exported Function
func Min_one_summit(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_valley : Exported Function
func Min_one_valley(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_one_zigzag : Exported Function
func Min_one_zigzag(data []float64) float64 {
	C := 1.0
	D := 1.0
	R := 1.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'e'
				} else if currentState == 'e' {
					D = 1.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'c' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'f' {
					C = 1.0
					D = 1.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_bump_on_decreasing_sequence : Exported Function
func Min_width_bump_on_decreasing_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_decreasing : Exported Function
func Min_width_decreasing(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
					D = 0.0
					R = math.Min(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
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
	return math.Min(R, C)
}

// Min_width_decreasing_sequence : Exported Function
func Min_width_decreasing_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 't' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_decreasing_terrace : Exported Function
func Min_width_decreasing_terrace(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_dip_on_increasing_sequence : Exported Function
func Min_width_dip_on_increasing_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_gorge : Exported Function
func Min_width_gorge(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_increasing : Exported Function
func Min_width_increasing(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Min(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return math.Min(R, C)
}

// Min_width_increasing_sequence : Exported Function
func Min_width_increasing_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_increasing_terrace : Exported Function
func Min_width_increasing_terrace(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_inflexion : Exported Function
func Min_width_inflexion(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_peak : Exported Function
func Min_width_peak(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_plain : Exported Function
func Min_width_plain(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_plateau : Exported Function
func Min_width_plateau(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_proper_plain : Exported Function
func Min_width_proper_plain(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_proper_plateau : Exported Function
func Min_width_proper_plateau(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_steady : Exported Function
func Min_width_steady(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Min(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_steady_sequence : Exported Function
func Min_width_steady_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_strictly_decreasing_sequence : Exported Function
func Min_width_strictly_decreasing_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_strictly_increasing_sequence : Exported Function
func Min_width_strictly_increasing_sequence(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_summit : Exported Function
func Min_width_summit(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_valley : Exported Function
func Min_width_valley(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
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
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_width_zigzag : Exported Function
func Min_width_zigzag(data []float64) float64 {
	C := float64(len(data))
	D := 0.0
	R := float64(len(data))
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = add(Dtemp, 1.0)
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = add(Dtemp, 1.0)
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'c' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'f' {
					C = float64(len(data))
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_bump_on_decreasing_sequence : Exported Function
func Min_surface_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_decreasing : Exported Function
func Min_surface_decreasing(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
					D = 0.0
					R = math.Min(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
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
	return math.Min(R, C)
}

// Min_surface_decreasing_sequence : Exported Function
func Min_surface_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_decreasing_terrace : Exported Function
func Min_surface_decreasing_terrace(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_dip_on_increasing_sequence : Exported Function
func Min_surface_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_gorge : Exported Function
func Min_surface_gorge(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_increasing : Exported Function
func Min_surface_increasing(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Min(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return math.Min(R, C)
}

// Min_surface_increasing_sequence : Exported Function
func Min_surface_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_increasing_terrace : Exported Function
func Min_surface_increasing_terrace(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_inflexion : Exported Function
func Min_surface_inflexion(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_peak : Exported Function
func Min_surface_peak(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_plain : Exported Function
func Min_surface_plain(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_plateau : Exported Function
func Min_surface_plateau(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_proper_plain : Exported Function
func Min_surface_proper_plain(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_proper_plateau : Exported Function
func Min_surface_proper_plateau(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = math.Min(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_steady : Exported Function
func Min_surface_steady(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = math.Min(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_steady_sequence : Exported Function
func Min_surface_steady_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_strictly_decreasing_sequence : Exported Function
func Min_surface_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_strictly_increasing_sequence : Exported Function
func Min_surface_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_summit : Exported Function
func Min_surface_summit(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_valley : Exported Function
func Min_surface_valley(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_surface_zigzag : Exported Function
func Min_surface_zigzag(data []float64) float64 {
	C := math.Inf(1)
	D := 0.0
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = add(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = add(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(1)
					D = 0.0
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_bump_on_decreasing_sequence : Exported Function
func Min_max_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_decreasing : Exported Function
func Min_max_decreasing(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
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
	return math.Min(R, C)
}

// Min_max_decreasing_sequence : Exported Function
func Min_max_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_decreasing_terrace : Exported Function
func Min_max_decreasing_terrace(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_dip_on_increasing_sequence : Exported Function
func Min_max_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_gorge : Exported Function
func Min_max_gorge(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_increasing : Exported Function
func Min_max_increasing(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return math.Min(R, C)
}

// Min_max_increasing_sequence : Exported Function
func Min_max_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_increasing_terrace : Exported Function
func Min_max_increasing_terrace(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_inflexion : Exported Function
func Min_max_inflexion(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_peak : Exported Function
func Min_max_peak(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_plain : Exported Function
func Min_max_plain(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_plateau : Exported Function
func Min_max_plateau(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_proper_plain : Exported Function
func Min_max_proper_plain(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_proper_plateau : Exported Function
func Min_max_proper_plateau(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_steady : Exported Function
func Min_max_steady(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = math.Inf(-1)
					R = math.Min(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_steady_sequence : Exported Function
func Min_max_steady_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_strictly_decreasing_sequence : Exported Function
func Min_max_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_strictly_increasing_sequence : Exported Function
func Min_max_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_summit : Exported Function
func Min_max_summit(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_valley : Exported Function
func Min_max_valley(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_max_zigzag : Exported Function
func Min_max_zigzag(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(-1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = math.Inf(-1)
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(1)
					D = math.Inf(-1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_bump_on_decreasing_sequence : Exported Function
func Min_min_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_decreasing : Exported Function
func Min_min_decreasing(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
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
	return math.Min(R, C)
}

// Min_min_decreasing_sequence : Exported Function
func Min_min_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_decreasing_terrace : Exported Function
func Min_min_decreasing_terrace(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_dip_on_increasing_sequence : Exported Function
func Min_min_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_gorge : Exported Function
func Min_min_gorge(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_increasing : Exported Function
func Min_min_increasing(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return math.Min(R, C)
}

// Min_min_increasing_sequence : Exported Function
func Min_min_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_increasing_terrace : Exported Function
func Min_min_increasing_terrace(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_inflexion : Exported Function
func Min_min_inflexion(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_peak : Exported Function
func Min_min_peak(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_plain : Exported Function
func Min_min_plain(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_plateau : Exported Function
func Min_min_plateau(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_proper_plain : Exported Function
func Min_min_proper_plain(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_proper_plateau : Exported Function
func Min_min_proper_plateau(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_steady : Exported Function
func Min_min_steady(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = math.Inf(1)
					R = math.Min(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_steady_sequence : Exported Function
func Min_min_steady_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_strictly_decreasing_sequence : Exported Function
func Min_min_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_strictly_increasing_sequence : Exported Function
func Min_min_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_summit : Exported Function
func Min_min_summit(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_valley : Exported Function
func Min_min_valley(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
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
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_min_zigzag : Exported Function
func Min_min_zigzag(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = math.Inf(1)
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(1)
					D = math.Inf(1)
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_bump_on_decreasing_sequence : Exported Function
func Min_range_bump_on_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_decreasing : Exported Function
func Min_range_decreasing(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
					R = math.Min(Rtemp, diff(data[i-1], data[i])) // R, found_e a0, Range Update
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
	return math.Min(R, C)
}

// Min_range_decreasing_sequence : Exported Function
func Min_range_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
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
	return math.Min(R, C)
}

// Min_range_decreasing_terrace : Exported Function
func Min_range_decreasing_terrace(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_dip_on_increasing_sequence : Exported Function
func Min_range_dip_on_increasing_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_gorge : Exported Function
func Min_range_gorge(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_increasing : Exported Function
func Min_range_increasing(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0                                       //neutral_f
					R = math.Min(Rtemp, diff(data[i-1], data[i])) // R, found_e a0, Range Update
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return math.Min(R, C)
}

// Min_range_increasing_sequence : Exported Function
func Min_range_increasing_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	H := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
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
	return math.Min(R, C)
}

// Min_range_increasing_terrace : Exported Function
func Min_range_increasing_terrace(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_inflexion : Exported Function
func Min_range_inflexion(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_peak : Exported Function
func Min_range_peak(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_plain : Exported Function
func Min_range_plain(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_plateau : Exported Function
func Min_range_plateau(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_proper_plain : Exported Function
func Min_range_proper_plain(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_proper_plateau : Exported Function
func Min_range_proper_plateau(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                     //neutral_f
					R = math.Min(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_steady : Exported Function
func Min_range_steady(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0                                                    //neutral_f
					R = math.Min(Rtemp, diff(diff(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_steady_sequence : Exported Function
func Min_range_steady_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(Ctemp, diff(Dtemp, data[i])) // C, in a0
					D = 0.0                               //neutral_f
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_strictly_decreasing_sequence : Exported Function
func Min_range_strictly_decreasing_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				H = data[i-1]
				H = math.Max(H, Htemp) // Holding onto the largest value for sequence
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				H = 0.0
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_strictly_increasing_sequence : Exported Function
func Min_range_strictly_increasing_sequence(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	H := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_summit : Exported Function
func Min_range_summit(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_valley : Exported Function
func Min_range_valley(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
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
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Min_range_zigzag : Exported Function
func Min_range_zigzag(data []float64) float64 {
	C := math.Inf(1) //max_f
	D := 0.0         //neutral_f
	R := math.Inf(1) //max_f
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'c'
				} else if currentState == 'c' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = diff(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0 //neutral_f
					currentState = 'a'
				} else if currentState == 'f' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = diff(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'f'
				} else if currentState == 'f' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'c' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'f' {
					C = math.Inf(1) //max_f
					D = 0.0         //neutral_f
					R = math.Min(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return math.Min(R, C)
}

// Sum_one_bump_on_decreasing_sequence : Exported Function
func Sum_one_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_decreasing : Exported Function
func Sum_one_decreasing(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
					D = 1.0
					R = add(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
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
	return add(R, C)
}

// Sum_one_decreasing_sequence : Exported Function
func Sum_one_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 't' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_decreasing_terrace : Exported Function
func Sum_one_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_dip_on_increasing_sequence : Exported Function
func Sum_one_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 1.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_gorge : Exported Function
func Sum_one_gorge(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_increasing : Exported Function
func Sum_one_increasing(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 1.0
					R = add(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return add(R, C)
}

// Sum_one_increasing_sequence : Exported Function
func Sum_one_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_increasing_terrace : Exported Function
func Sum_one_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_inflexion : Exported Function
func Sum_one_inflexion(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_peak : Exported Function
func Sum_one_peak(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_plain : Exported Function
func Sum_one_plain(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_plateau : Exported Function
func Sum_one_plateau(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_proper_plain : Exported Function
func Sum_one_proper_plain(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_proper_plateau : Exported Function
func Sum_one_proper_plateau(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 1.0
					R = add(Rtemp, math.Max(Dtemp, 0.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_steady : Exported Function
func Sum_one_steady(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 1.0
					R = add(Rtemp, math.Max(math.Max(Dtemp, 0.0), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_steady_sequence : Exported Function
func Sum_one_steady_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_strictly_decreasing_sequence : Exported Function
func Sum_one_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_strictly_increasing_sequence : Exported Function
func Sum_one_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, 0.0), data[i]) // C, found a0
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = 1.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_summit : Exported Function
func Sum_one_summit(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 1.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_valley : Exported Function
func Sum_one_valley(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, 0.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_one_zigzag : Exported Function
func Sum_one_zigzag(data []float64) float64 {
	C := 0.0
	D := 1.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'e'
				} else if currentState == 'e' {
					D = 1.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Max(Dtemp, 0.0)
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Max(Ctemp, math.Max(Dtemp, 0.0)) // C, in a1
					D = 1.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Max(Dtemp, 0.0) // C, found a1
					D = 1.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 1.0
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = 1.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_bump_on_decreasing_sequence : Exported Function
func Sum_width_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_decreasing : Exported Function
func Sum_width_decreasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
					D = 0.0
					R = add(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
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
	return add(R, C)
}

// Sum_width_decreasing_sequence : Exported Function
func Sum_width_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_decreasing_terrace : Exported Function
func Sum_width_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_dip_on_increasing_sequence : Exported Function
func Sum_width_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_gorge : Exported Function
func Sum_width_gorge(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_increasing : Exported Function
func Sum_width_increasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = add(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return add(R, C)
}

// Sum_width_increasing_sequence : Exported Function
func Sum_width_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_increasing_terrace : Exported Function
func Sum_width_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_inflexion : Exported Function
func Sum_width_inflexion(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_peak : Exported Function
func Sum_width_peak(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_plain : Exported Function
func Sum_width_plain(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_plateau : Exported Function
func Sum_width_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_proper_plain : Exported Function
func Sum_width_proper_plain(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_proper_plateau : Exported Function
func Sum_width_proper_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, 1.0)) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_steady : Exported Function
func Sum_width_steady(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = add(Rtemp, add(add(Dtemp, 1.0), 1.0)) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_steady_sequence : Exported Function
func Sum_width_steady_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_strictly_decreasing_sequence : Exported Function
func Sum_width_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_strictly_increasing_sequence : Exported Function
func Sum_width_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, 1.0), 1.0) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_summit : Exported Function
func Sum_width_summit(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, 1.0)
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_valley : Exported Function
func Sum_width_valley(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, 1.0)
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, 1.0)
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_width_zigzag : Exported Function
func Sum_width_zigzag(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = add(Dtemp, 1.0)
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = add(Dtemp, 1.0)
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = add(Ctemp, add(Dtemp, 1.0)) // C, in a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = add(Dtemp, 1.0) // C, found a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_bump_on_decreasing_sequence : Exported Function
func Sum_surface_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_decreasing : Exported Function
func Sum_surface_decreasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
					D = 0.0
					R = add(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
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
	return add(R, C)
}

// Sum_surface_decreasing_sequence : Exported Function
func Sum_surface_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_decreasing_terrace : Exported Function
func Sum_surface_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_dip_on_increasing_sequence : Exported Function
func Sum_surface_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_gorge : Exported Function
func Sum_surface_gorge(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_increasing : Exported Function
func Sum_surface_increasing(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = add(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return add(R, C)
}

// Sum_surface_increasing_sequence : Exported Function
func Sum_surface_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = add(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_increasing_terrace : Exported Function
func Sum_surface_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_inflexion : Exported Function
func Sum_surface_inflexion(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_peak : Exported Function
func Sum_surface_peak(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_plain : Exported Function
func Sum_surface_plain(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_plateau : Exported Function
func Sum_surface_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_proper_plain : Exported Function
func Sum_surface_proper_plain(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_proper_plateau : Exported Function
func Sum_surface_proper_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0
					R = add(Rtemp, add(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_steady : Exported Function
func Sum_surface_steady(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0
					R = add(Rtemp, add(add(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_steady_sequence : Exported Function
func Sum_surface_steady_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_strictly_decreasing_sequence : Exported Function
func Sum_surface_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_strictly_increasing_sequence : Exported Function
func Sum_surface_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(add(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					C = add(Ctemp, add(Dtemp, data[i])) // C, in a0
					D = 0.0
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_summit : Exported Function
func Sum_surface_summit(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = add(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_valley : Exported Function
func Sum_surface_valley(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
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
				} else if currentState == 'r' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 't'
				} else if currentState == 't' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = add(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = add(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_surface_zigzag : Exported Function
func Sum_surface_zigzag(data []float64) float64 {
	C := 0.0
	D := 0.0
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = add(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 'a'
				} else if currentState == 'f' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = add(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = add(Ctemp, add(Dtemp, data[i-1])) // C, in a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = add(Dtemp, data[i-1]) // C, found a1
					D = 0.0
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_bump_on_decreasing_sequence : Exported Function
func Sum_max_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_decreasing : Exported Function
func Sum_max_decreasing(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
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
	return add(R, C)
}

// Sum_max_decreasing_sequence : Exported Function
func Sum_max_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_decreasing_terrace : Exported Function
func Sum_max_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_dip_on_increasing_sequence : Exported Function
func Sum_max_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(-1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_gorge : Exported Function
func Sum_max_gorge(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_increasing : Exported Function
func Sum_max_increasing(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return add(R, C)
}

// Sum_max_increasing_sequence : Exported Function
func Sum_max_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_increasing_terrace : Exported Function
func Sum_max_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_inflexion : Exported Function
func Sum_max_inflexion(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_peak : Exported Function
func Sum_max_peak(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_plain : Exported Function
func Sum_max_plain(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_plateau : Exported Function
func Sum_max_plateau(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_proper_plain : Exported Function
func Sum_max_proper_plain(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_proper_plateau : Exported Function
func Sum_max_proper_plateau(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_steady : Exported Function
func Sum_max_steady(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = math.Inf(-1)
					R = add(Rtemp, math.Max(math.Max(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_steady_sequence : Exported Function
func Sum_max_steady_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_strictly_decreasing_sequence : Exported Function
func Sum_max_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_strictly_increasing_sequence : Exported Function
func Sum_max_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(math.Max(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i])) // C, in a0
					D = math.Inf(-1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_summit : Exported Function
func Sum_max_summit(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(-1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_valley : Exported Function
func Sum_max_valley(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_max_zigzag : Exported Function
func Sum_max_zigzag(data []float64) float64 {
	C := 0.0
	D := math.Inf(-1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = math.Inf(-1)
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Max(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Max(Ctemp, math.Max(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(-1)
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Max(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(-1)
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = math.Inf(-1)
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = math.Inf(-1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_bump_on_decreasing_sequence : Exported Function
func Sum_min_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_decreasing : Exported Function
func Sum_min_decreasing(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
					D = math.Inf(1)
					R = add(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
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
	return add(R, C)
}

// Sum_min_decreasing_sequence : Exported Function
func Sum_min_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_decreasing_terrace : Exported Function
func Sum_min_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_dip_on_increasing_sequence : Exported Function
func Sum_min_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'v' {
					D = math.Inf(1)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_gorge : Exported Function
func Sum_min_gorge(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_increasing : Exported Function
func Sum_min_increasing(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return add(R, C)
}

// Sum_min_increasing_sequence : Exported Function
func Sum_min_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_increasing_terrace : Exported Function
func Sum_min_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_inflexion : Exported Function
func Sum_min_inflexion(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_peak : Exported Function
func Sum_min_peak(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_plain : Exported Function
func Sum_min_plain(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_plateau : Exported Function
func Sum_min_plateau(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_proper_plain : Exported Function
func Sum_min_proper_plain(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_proper_plateau : Exported Function
func Sum_min_proper_plateau(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_steady : Exported Function
func Sum_min_steady(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = math.Inf(1)
					R = add(Rtemp, math.Min(math.Min(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_steady_sequence : Exported Function
func Sum_min_steady_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_strictly_decreasing_sequence : Exported Function
func Sum_min_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_strictly_increasing_sequence : Exported Function
func Sum_min_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i]) // C, found a0
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i])) // C, in a0
					D = math.Inf(1)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_summit : Exported Function
func Sum_min_summit(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Inf(1)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_valley : Exported Function
func Sum_min_valley(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
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
				} else if currentState == 'r' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_min_zigzag : Exported Function
func Sum_min_zigzag(data []float64) float64 {
	C := 0.0
	D := math.Inf(1)
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = math.Inf(1)
					currentState = 'a'
				} else if currentState == 'f' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = math.Min(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i-1])) // C, in a1
					D = math.Inf(1)
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = math.Min(Dtemp, data[i-1]) // C, found a1
					D = math.Inf(1)
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = math.Inf(1)
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = math.Inf(1)
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_bump_on_decreasing_sequence : Exported Function
func Sum_range_bump_on_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_decreasing : Exported Function
func Sum_range_decreasing(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
					D = 0.0                                  //neutral_f
					R = add(Rtemp, diff(data[i-1], data[i])) // R, found_e a0, Range Update
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
	return add(R, C)
}

// Sum_range_decreasing_sequence : Exported Function
func Sum_range_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
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
	return add(R, C)
}

// Sum_range_decreasing_terrace : Exported Function
func Sum_range_decreasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_dip_on_increasing_sequence : Exported Function
func Sum_range_dip_on_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 't'
				} else if currentState == 't' {
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'v'
				} else if currentState == 'v' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					currentState = 's'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'v' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_gorge : Exported Function
func Sum_range_gorge(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_increasing : Exported Function
func Sum_range_increasing(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					D = 0.0                                  //neutral_f
					R = add(Rtemp, diff(data[i-1], data[i])) // R, found_e a0, Range Update
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
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
	return add(R, C)
}

// Sum_range_increasing_sequence : Exported Function
func Sum_range_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	H := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
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
	return add(R, C)
}

// Sum_range_increasing_terrace : Exported Function
func Sum_range_increasing_terrace(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_inflexion : Exported Function
func Sum_range_inflexion(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 't'
				} else if currentState == 'r' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_peak : Exported Function
func Sum_range_peak(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_plain : Exported Function
func Sum_range_plain(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_plateau : Exported Function
func Sum_range_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_proper_plain : Exported Function
func Sum_range_proper_plain(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_proper_plateau : Exported Function
func Sum_range_proper_plateau(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					currentState = 'r'
				} else if currentState == 't' {
					D = 0.0 //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					currentState = 's'
				} else if currentState == 't' {
					D = 0.0                                //neutral_f
					R = add(Rtemp, diff(Dtemp, data[i-1])) // R, found_e a1
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_steady : Exported Function
func Sum_range_steady(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					D = 0.0                                               //neutral_f
					R = add(Rtemp, diff(diff(Dtemp, data[i-1]), data[i])) // R, found_e a0
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_steady_sequence : Exported Function
func Sum_range_steady_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(Ctemp, diff(Dtemp, data[i])) // C, in a0
					D = 0.0                               //neutral_f
					currentState = 'r'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_strictly_decreasing_sequence : Exported Function
func Sum_range_strictly_decreasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] < data[i-1] {
				H = data[i-1]
				H = math.Max(H, Htemp) // Holding onto the largest value for sequence
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				H = 0.0
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_strictly_increasing_sequence : Exported Function
func Sum_range_strictly_increasing_sequence(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	H := math.Inf(1)
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			Htemp := float64(H)
			if data[i] > data[i-1] {
				H = data[i-1]
				H = math.Min(H, Htemp)
				if currentState == 's' {
					C = diff(diff(Dtemp, data[i-1]), data[i]) // C, found a0
					D = 0.0                                   //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					C = diff(H, data[i]) // C, in a0
					D = 0.0              //neutral_f
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				H = math.Inf(1)
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_summit : Exported Function
func Sum_range_summit(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				} else if currentState == 'u' {
					D = 0.0 //neutral_f
					currentState = 's'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				} else if currentState == 'u' {
					D = diff(Dtemp, data[i-1])
					currentState = 'u'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_valley : Exported Function
func Sum_range_valley(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
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
				} else if currentState == 'r' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 't'
				} else if currentState == 't' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 't'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'r'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 'r'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'r' {
					D = diff(Dtemp, data[i-1])
					currentState = 'r'
				} else if currentState == 't' {
					D = diff(Dtemp, data[i-1])
					currentState = 't'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}

// Sum_range_zigzag : Exported Function
func Sum_range_zigzag(data []float64) float64 {
	C := 0.0
	D := 0.0 //neutral_f
	R := 0.0
	currentState := 's'
	DataLen := len(data)
	for i := 1; i < DataLen; i++ {
		if i < DataLen {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					currentState = 'a'
				} else if currentState == 'a' {
					currentState = 'a'
				} else if currentState == 'b' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'c'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 'a'
				} else if currentState == 'd' {
					D = diff(Dtemp, data[i-1])
					currentState = 'e'
				} else if currentState == 'e' {
					D = 0.0 //neutral_f
					currentState = 'a'
				} else if currentState == 'f' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 'c'
				}
			} else if data[i] < data[i-1] {
				if currentState == 's' {
					currentState = 'd'
				} else if currentState == 'a' {
					D = diff(Dtemp, data[i-1])
					currentState = 'b'
				} else if currentState == 'b' {
					currentState = 'd'
				} else if currentState == 'c' {
					C = diff(Ctemp, diff(Dtemp, data[i-1])) // C, in a1
					D = 0.0                                 //neutral_f
					currentState = 'f'
				} else if currentState == 'd' {
					currentState = 'd'
				} else if currentState == 'e' {
					C = diff(Dtemp, data[i-1]) // C, found a1
					D = 0.0                    //neutral_f
					currentState = 'f'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 'd'
				}
			} else if data[i] == data[i-1] {
				if currentState == 's' {
					currentState = 's'
				} else if currentState == 'a' {
					currentState = 's'
				} else if currentState == 'b' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'c' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				} else if currentState == 'd' {
					currentState = 's'
				} else if currentState == 'e' {
					D = 0.0 //neutral_f
					currentState = 's'
				} else if currentState == 'f' {
					C = 0.0
					D = 0.0 //neutral_f
					R = add(Rtemp, Ctemp)
					currentState = 's'
				}
			}
			_ = Ctemp // Temporary fix
			_ = Dtemp // Temporary fix
			_ = Rtemp // Temporary fix
		}
	}
	return add(R, C)
}
