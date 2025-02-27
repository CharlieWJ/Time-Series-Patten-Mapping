# Time series pattern recognition using generated Go code

Generated code written in GoLang that focuses on analysing and extracting patterns from time series.

The generated functions are based on a model presented in the second volume of the [Global Constraint Catalog](https://arxiv.org/abs/1609.08925).

The original library of functions was generated in Python, but did not support the usage of 'range' functions. ([Original](https://github.com/allarddenis/time-series-pattern-recognition)) 


The model uses accumulators to extract aggregations on features of a pattern occurrences in a time series.


# Summary
1. [Models](#models)
2. [How to](#how-to)
3. [Side Notes](#Side-Notes)

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

Note: The ‘range’ feature is only used with monotonic patterns (i.e. increasing/decreasing(_sequence), strictly_increasing/decreasing_sequence), but still generates functions for constant patterns (see page 12 of the catalog).

# How to

## How to generate functions

From **/src** execute the following command :

```
python generator.py    
```

This should generate **generatedInGo.go** file in **/src/generatedInGo**.

## How to test
---
## 1 ) Running all tests
There are currently 150 test cases provided. You can find them in one of two places: 
- The first being inside of **/src/inputs/input_go_tests.go** 
- and the other in **/src/testy.go**

The test cases found in both *input_go_tests.go* and _testy.go_ are the same, but if you wish to add more cases they must be added to the _**Tester**_ struct in order to be included when running all cases. When adding new tests be sure to include "*generated.*" before the function name.

To run all of the tests navigate to **/src** and execute the following command:
```
go run testy.go
```


## 2 ) Standalone testing
From **/src** execute the following command :

```
go run main.go    
```

This should output test results in command window.

**Note :** Currently, only 149 out of the 396 functions generated have tests.


**Tested functions** : 149 out of 396 generated

## How to use

Choose your desired function from **/src/generated/generatedInGo.go**.

All generated functions are **standalone**; meaning you only need to copy-paste them to use it. Remember to **import: "fmt" & "math"** in your file as many of the functions utilize **math.Inf/Max/Min**. Also, be sure to include an 'add()' method as many of the functions require one :
```go
func add(x float64, y float64) float64{
    return (x+y)
}
```

Data inputs are expected to be an integer array/slice :

```go
data = []float64{0,1,2,3,4,5,6,7,8,9}
```

Here is an example of a '.go' file using only **Min_min_increasing_sequence** function :

```go
package main 

import(
    "fmt"
    "math"
)

func Min_min_increasing_sequence(data []float64) float64 {
	C := math.Inf(1)
	D := math.Inf(1)
	R := math.Inf(1)
	currentState := 's'
	for i := 1; i < len(data); i++ {
		if i < len(data) {
			Ctemp := float64(C)
			Dtemp := float64(D)
			Rtemp := float64(R)
			if data[i] > data[i-1] {
				if currentState == 's' {
					C = math.Min(math.Min(Dtemp, data[i-1]), data[i])
					D = math.Inf(1)
					currentState = 't'
				} else if currentState == 't' {
					C = math.Min(Ctemp, math.Min(Dtemp, data[i]))
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
			_ = Ctemp
			_ = Dtemp
			_ = Rtemp
		}
	}
	return math.Min(R, C)
}



func main(){
    data := []float64{4,3,5,5,2,1,1,3,3,4,6,6,3,1,3,3}
    fmt.Println(Min_min_increasing_sequence(data))
    // Should print '1'
    // Can run from the terminal line using: 
    //          $ go run <File name here>.go
}

```

## Side Notes
UPDATED EXTRA FUNCTION LIST:
All Max/Min_one_ funcs (44)


The generator generates approximately 165 extra functions, being:
- 44 '\_one\_' functions that are not tested or found in the catalog (sum_one is denoted as NB in the catalog, this requires attention. There are 22 sum_one funcs that need tests)
- 48 '\_range\_' functions, as only montonic patterns are tested (Fixed)
- 30 '\_max\_' functions (_max_gorge/(dec/inc_terrace)/)
- 15 '\_width\_' functions (Max/Min/Sum_width_(dip/bump/steady/inc/dec) (5 each -> 15))
- 6 '\_min/height\_' functions (Max/Min/Sum_min_(summit/peak))
- 5 'Sum_' funcs (Sum_width_dip/bump/steady/increasing/decreasing) (probably applies to min and max too)

It is important to note that functions that should be labled as '<something\>\_height\_<something\>' were just named with '\_min\_'. This is because in the Global Constraint Catalog the ‘min’ feature is renamed ‘height’ when used together with constant patterns, however that was not something accounted for when generating functions. I may fix this once everything is accounted for.

In total, there are 231 functions that I am certain about which work. Those 231 functions have all been tested with the examples that are provided in the catalog and can be found in the _testy.go_ script. 

There are also a total of 21 different test cases used, they can be found in the 'res' folder. There is also a test case legend markdown file in the 'res' folder to see which cases correspond to which functions. The functions that are not generated are NB_<something\>, which there are 21 of.