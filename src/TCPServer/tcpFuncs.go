package tcpserver

import (
	"fmt"
	"os"

	gen "../generatedInGo"
)

// Note to self: I might be using a lot more memory than I should be, think of ways to optimize.
type results struct {
	name     string
	expected string
	result   string
}

// WriteToFile : Makes file with the results
func WriteToFile(testcase string) {
	f, err := os.Create("../res/results/" + testcase + ".md")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("# ||" + testcase + "||\nFunction | Result | Expected\n---|---|---\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	temp := runtests(testcase)
	length := len(temp)
	for i := 0; i < length; i++ {
		e := "|" + temp[i].name + "| Result: " + temp[i].result + "| Expected: ~" + temp[i].expected
		fmt.Fprintln(f, e)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// This might be my longest function in this entire repo... don't know yet though
func runtests(testcase string) []results {
	holder := []results{}
	if testcase == "CASE0" {
		length := len(gen.Case0)
		for i := 0; i < length; i++ {
			x := gen.Case0[i]
			x.Args.Data = Sens1Prop1
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
			//fmt.Println(holder[i])
		}
	} else if testcase == "CASE1" {
		length := len(gen.Case1)
		for i := 0; i < length; i++ {
			x := gen.Case1[i]
			x.Args.Data = Sens1Prop2
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
			//fmt.Println(holder[i])
		}
	} else if testcase == "CASE2" {
		length := len(gen.Case2)
		for i := 0; i < length; i++ {
			x := gen.Case2[i]
			x.Args.Data = Sens1Prop3
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
			//fmt.Println(holder[i])
		}
	} else if testcase == "CASE3" {
		length := len(gen.Case3)
		for i := 0; i < length; i++ {
			x := gen.Case2[i]
			x.Args.Data = Sens1Prop4
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
			//fmt.Println(holder[i])
		}
	} else if testcase == "CASE4" {
		length := len(gen.Case4)
		for i := 0; i < length; i++ {
			x := gen.Case2[i]
			x.Args.Data = Sens1Prop5
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
			//fmt.Println(holder[i])
		}
	} else if testcase == "CASE5" {
		length := len(gen.Case5)
		for i := 0; i < length; i++ {
			x := gen.Case2[i]
			x.Args.Data = Sens1Prop6
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
			//fmt.Println(holder[i])
		}
	}
	return holder
}
