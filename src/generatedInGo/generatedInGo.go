// ----------------------------------------------------------------------------
// This file was auto-generated on 2020-06-11
// By Charles W. Jeffries using an altered Python Script.
// The original script was provided by: Florine Cercle & Denis Allard
// Original Source Code : https://github.com/allarddenis/time-series-pattern-recognition
// ----------------------------------------------------------------------------

package generatedInGo

import(
	"fmt"
	"math"
)

func add(x float64, y float64) float64 { return (x+y) }
type Tuple struct { a,b interface{} }

func GetMax(x float64, y float64) float64 {
    if x>=y { 
        return x
    } else { return y }
}
// Currently not using this method

func GetMin(x float64, y float64) float64 {
    if x>=y { 
        return y
    } else { return x }
}
// Currently not using this method


func Max_one_bump_on_decreasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_decreasing(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_decreasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_decreasing_terrace(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_dip_on_increasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_gorge(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_increasing(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_increasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_increasing_terrace(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_inflexion(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_peak(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_plain(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_plateau(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_proper_plain(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_proper_plateau(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_steady(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_steady_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_strictly_decreasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_strictly_increasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_summit(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_valley(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_one_zigzag(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 1.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
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
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 1.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_decreasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(add(D_temp,1.0),1))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_gorge(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_increasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(add(D_temp,1.0),1))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_increasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_inflexion(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_peak(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_proper_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_proper_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_steady(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(add(D_temp,1.0),1))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_steady_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_summit(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_valley(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_width_zigzag(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
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
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_decreasing(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_decreasing_terrace(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_gorge(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_increasing(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_increasing_terrace(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_inflexion(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_peak(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_plain(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_plateau(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_proper_plain(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_proper_plateau(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_steady(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_steady_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_summit(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_valley(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_surface_zigzag(data []float64) float64{
    C := math.Inf(-1)
    D := 0.0
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
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
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(-1)                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_decreasing(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_decreasing_terrace(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_gorge(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_increasing(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_increasing_terrace(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_inflexion(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_peak(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_plain(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_plateau(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_proper_plain(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_proper_plateau(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_steady(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_steady_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_summit(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_valley(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_max_zigzag(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(-1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = math.Inf(-1)                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
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
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = math.Inf(-1)                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(-1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_decreasing(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_decreasing_terrace(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_gorge(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_increasing(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_increasing_terrace(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_inflexion(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_peak(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_plain(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_plateau(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_proper_plain(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_proper_plateau(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_steady(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_steady_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_summit(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_valley(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_min_zigzag(data []float64) float64{
    C := math.Inf(-1)
    D := math.Inf(1)
    R := math.Inf(-1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = math.Inf(1)                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
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
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = math.Inf(1)                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(-1)                    
                    D = math.Inf(1)                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_decreasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Max(R_temp,((D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_gorge(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_increasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Max(R_temp,((D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_increasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_inflexion(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_peak(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_proper_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_proper_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_steady(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Max(R_temp,((D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_steady_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_summit(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_valley(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Max_range_zigzag(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
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
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = math.Max(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Max(R,C)    
}
func Min_one_bump_on_decreasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_decreasing(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_decreasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_decreasing_terrace(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_dip_on_increasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_gorge(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_increasing(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_increasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_increasing_terrace(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_inflexion(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_peak(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_plain(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_plateau(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_proper_plain(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_proper_plateau(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_steady(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_steady_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_strictly_decreasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_strictly_increasing_sequence(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_summit(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_valley(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_one_zigzag(data []float64) float64{
    C := 1.0
    D := 1.0
    R := 1.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 1.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
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
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 1.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 1.0                    
                    D = 1.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_bump_on_decreasing_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_decreasing(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(add(D_temp,1.0),1))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_decreasing_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_decreasing_terrace(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_dip_on_increasing_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_gorge(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_increasing(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(add(D_temp,1.0),1))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_increasing_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_increasing_terrace(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_inflexion(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_peak(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_plain(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_plateau(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_proper_plain(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_proper_plateau(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_steady(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(add(D_temp,1.0),1))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_steady_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_strictly_decreasing_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_strictly_increasing_sequence(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_summit(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_valley(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_width_zigzag(data []float64) float64{
    C := len(data)
    D := 0.0
    R := len(data)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
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
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = len(data)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_decreasing(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_decreasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_gorge(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_increasing(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_increasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_inflexion(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_peak(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_plain(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_plateau(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_proper_plain(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_proper_plateau(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_steady(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_steady_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_summit(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_valley(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_surface_zigzag(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
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
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_decreasing(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_decreasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_gorge(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_increasing(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_increasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_inflexion(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_peak(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_plain(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_plateau(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_proper_plain(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_proper_plateau(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_steady(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_steady_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_summit(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_valley(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_max_zigzag(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(-1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = math.Inf(-1)                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
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
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = math.Inf(-1)                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = math.Inf(-1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_decreasing(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_decreasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_gorge(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_increasing(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_increasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_inflexion(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_peak(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_plain(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_plateau(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_proper_plain(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_proper_plateau(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_steady(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_steady_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_summit(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_valley(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_min_zigzag(data []float64) float64{
    C := math.Inf(1)
    D := math.Inf(1)
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = math.Inf(1)                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
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
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = math.Inf(1)                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = math.Inf(1)                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_bump_on_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_decreasing(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Min(R_temp,((D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_decreasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_dip_on_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_gorge(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_increasing(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = math.Min(R_temp,((D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_increasing_terrace(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_inflexion(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_peak(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_plain(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_plateau(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_proper_plain(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_proper_plateau(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_steady(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = math.Min(R_temp,((D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_steady_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_strictly_decreasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_strictly_increasing_sequence(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_summit(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_valley(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Min_range_zigzag(data []float64) float64{
    C := math.Inf(1)
    D := 0.0
    R := math.Inf(1)
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
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
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = math.Inf(1)                    
                    D = 0.0                    
                    R = math.Min(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return math.Min(R,C)    
}
func Sum_one_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_one_decreasing(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_one_gorge(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_increasing(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_one_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_increasing_terrace(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_inflexion(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_peak(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_plain(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_plateau(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_proper_plain(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 1.0                    
                    R = add(R_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_proper_plateau(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(D_temp,0.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_steady(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(math.Max(D_temp,0.0),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_steady_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,0.0),data[i])                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = 1.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_summit(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
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
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_valley(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_one_zigzag(data []float64) float64{
    C := 0.0
    D := 1.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 1.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Max(D_temp,0.0)                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Max(C_temp,math.Max(D_temp,0.0))                    
                    D = 1.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Max(D_temp,0.0)                    
                    D = 1.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
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
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 1.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 1.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_width_decreasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = add(R_temp,add(add(D_temp,1.0),1))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_width_gorge(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_increasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = add(R_temp,add(add(D_temp,1.0),1))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_width_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_increasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_inflexion(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_peak(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_proper_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_proper_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(D_temp,1.0))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_steady(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(add(D_temp,1.0),1))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_steady_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,1.0),1)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,1))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_summit(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
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
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_valley(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,1.0)                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_width_zigzag(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = add(D_temp,1.0)                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = add(C_temp,add(D_temp,1.0))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = add(D_temp,1.0)                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
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
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_surface_decreasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = add(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_surface_gorge(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_increasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = add(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_surface_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_increasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_inflexion(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_peak(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_proper_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_proper_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_steady(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,add(add(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_steady_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(add(D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = add(C_temp,add(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_summit(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
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
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_valley(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_surface_zigzag(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = add(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = add(C_temp,add(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = add(D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
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
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_max_decreasing(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_max_gorge(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_increasing(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_max_increasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_increasing_terrace(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_inflexion(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_peak(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_plain(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_plateau(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_proper_plain(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(-1)                    
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_proper_plateau(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_steady(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Max(math.Max(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_steady_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(math.Max(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i]))                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_summit(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
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
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_valley(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_max_zigzag(data []float64) float64{
    C := 0.0
    D := math.Inf(-1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = math.Inf(-1)                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Max(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Max(C_temp,math.Max(D_temp,data[i-1]))                    
                    D = math.Inf(-1)                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Max(D_temp,data[i-1])                    
                    D = math.Inf(-1)                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
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
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = math.Inf(-1)                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = math.Inf(-1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_min_decreasing(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_min_gorge(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_increasing(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_min_increasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_increasing_terrace(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_inflexion(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_peak(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_plain(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_plateau(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_proper_plain(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = math.Inf(1)                    
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_proper_plateau(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Min(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_steady(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,math.Min(math.Min(D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_steady_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(math.Min(D_temp,data[i-1]),data[i])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i]))                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_summit(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
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
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_valley(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_min_zigzag(data []float64) float64{
    C := 0.0
    D := math.Inf(1)
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = math.Inf(1)                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = math.Min(D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = math.Min(C_temp,math.Min(D_temp,data[i-1]))                    
                    D = math.Inf(1)                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = math.Min(D_temp,data[i-1])                    
                    D = math.Inf(1)                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
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
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = math.Inf(1)                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = math.Inf(1)                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_bump_on_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_range_decreasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = add(R_temp,((D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_decreasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_dip_on_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 't'                    
                } else if currentState == 't' {                
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'v'                    
                } else if currentState == 'v' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_range_gorge(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_increasing(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    D = 0.0                    
                    R = add(R_temp,((D_temp,data[i-1]),data[i]))                    
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
        }        
    }    
    return add(R,C)    
}
func Sum_range_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_increasing_terrace(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_inflexion(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 't'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_peak(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_proper_plain(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    currentState = 's'                    
                } else if currentState == 't' {                
                    D = 0.0                    
                    R = add(R_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_proper_plateau(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,(D_temp,data[i-1]))                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_steady(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
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
                    R = add(R_temp,((D_temp,data[i-1]),data[i]))                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_steady_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_strictly_decreasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_strictly_increasing_sequence(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = ((D_temp,data[i-1]),data[i])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    C = (C_temp,(D_temp,data[i]))                    
                    D = 0.0                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_summit(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
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
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                } else if currentState == 'u' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'u'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_valley(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 't'                    
                } else if currentState == 't' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 't'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'r'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'r'                    
                }                
            } else if data[i] == data[i-1] {            
                if currentState == 's' {                
                    currentState = 's'                    
                } else if currentState == 'r' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'r'                    
                } else if currentState == 't' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 't'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
func Sum_range_zigzag(data []float64) float64{
    C := 0.0
    D := 0.0
    R := 0.0
    currentState := 's'
    for i := 1; i < len(data); i++ {    
        if i<len(data) {        
            C_temp := C            
            D_temp := D            
            R_temp := R            
            if data[i] > data[i-1] {            
                if currentState == 's' {                
                    currentState = 'a'                    
                } else if currentState == 'a' {                
                    currentState = 'a'                    
                } else if currentState == 'b' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'c'                    
                } else if currentState == 'c' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 'a'                    
                } else if currentState == 'd' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'e'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 'a'                    
                } else if currentState == 'f' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'c'                    
                }                
            } else if data[i] < data[i-1] {            
                if currentState == 's' {                
                    currentState = 'd'                    
                } else if currentState == 'a' {                
                    D = (D_temp,data[i-1])                    
                    currentState = 'b'                    
                } else if currentState == 'b' {                
                    currentState = 'd'                    
                } else if currentState == 'c' {                
                    C = (C_temp,(D_temp,data[i-1]))                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'd' {                
                    currentState = 'd'                    
                } else if currentState == 'e' {                
                    C = (D_temp,data[i-1])                    
                    D = 0.0                    
                    currentState = 'f'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
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
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                } else if currentState == 'd' {                
                    currentState = 's'                    
                } else if currentState == 'e' {                
                    D = 0.0                    
                    currentState = 's'                    
                } else if currentState == 'f' {                
                    C = 0.0                    
                    D = 0.0                    
                    R = add(R_temp,C_temp)                    
                    currentState = 's'                    
                }                
            }            
        }        
    }    
    return add(R,C)    
}
