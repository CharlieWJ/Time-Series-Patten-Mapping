# Time series pattern recognition using generated Go code

Generated code written in Go that focuses on analysing and extracting patterns from time series.

The generated functions are based on a github repo that generated code with the same functionality, but in Python. (see [original](https://github.com/allarddenis/time-series-pattern-recognition)). 

The model uses accumulators to extract aggregations on features of a pattern occurrences in a time series.


# Summary
1. [Models](#models)
2. [How to](#how-to)
3. [Contributors](#contributors)

# Models

## Patterns

Pattern name | Regex
------------ | -------------
bump_on_decreasing_sequence | >><>>
decreasing | >
increasing | <
decreasing_sequence | >(>|=)*>|>
decreasing_terrace | >=+>
dip_on_increasing_sequence | <<><<
gorge | (>\|(>(=\|>)*>))(<\|(<(=\|<)*<))
increasing_sequence | <(<\|=)*<\|<
increasing_terrace | <=+<
inflexion | <(<\|=)*>\|>(>\|=)*<
peak | <(=\|<)*(>\|=)*>
plain | >=*<
plateau | <=*>
proper_plain | >=+<
proper_plateau | <=+>
steady | =
steady_sequence | =+
strictly_increasing_sequence | <+
strictly_decreasing_sequence | >+
summit | (<\|(<(=\|<)*<))(>\|(>(=\|>)*>))
valley | >(=\|>)*(<\|=)*<
zigzag | (<>)+(<\|<>)\|(><)+(>\|><)

Example :

```
Data : [4,4,2,2,3,5,5,6,3,1,1,2,2,2,2,2,2,1]
Signature : =>=<<=<>>=<=====>

Peak regex : <(=|<)*(>|=)*>

Peak pattern occurrences : <<=<>> & <=====>
```

## Features

Name | Meaning
------------ | -------------
one | Identity feature
width | Width of a pattern occurrence
height | Height of a pattern occurrence
max | Largest value among items of a pattern occurrence
min | Lowest value among items of a pattern occurrence
surface | Surface of a pattern occurrence

## Aggregators

Name | Meaning
------------ | -------------
max | Largest value among features of pattern occurrences
min | Lowest value among features of pattern occurrences
sum | Sum of feature values of pattern occurrences

# How to

## How to generate functions

From **/src** execute the following command :

```
python generator.py    
```

This should generate **generatedInGo.go** file in **/src/generated**.

## How to test

From **/src** execute the following command :

```
go run main.go    
```

This should output test results in command window.

**Note :** Currently, only 123 out of the 396 functions generated have tests.

Some of those without tests include: 
- all 'range' functions. 
- ~1/2 zig_zag functions


**Tested functions** : 123 out of 396 generated

## How to use

Choose your desired function from **/src/generated/generatedInGo.go**.

All generated functions are **standalone**; meaning you only need to copy-paste them to use it. Remember to **import: "fmt" & "math"** in your file as many of the functions utilize **math.Inf/Max/Min**. Also, be sure to include an 'add()' method as many of the functions require one :
```go
func add(x int, y int) int{
    return (x+y)
}
```

Data inputs are expected to be an integer array/slice :

```go
data = []float64{0,1,2,3,4,5,6,7,8,9}
```

Here is an example of a '.go' file using only **min_min_increasing_sequence** function :

```go
package main 

import(
    "fmt"
    "math"
)

func min_min_increasing_sequence(data []float64) float64{
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


func main(){
    data := []float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
    fmt.Println(min_min_increasing_sequence(data))
    // Should print '1'
    // Can run from the terminal line using: 
    //          $ go run <File name here>.go
}

```
