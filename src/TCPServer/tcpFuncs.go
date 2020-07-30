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
	//temp := runtests(testcase)
	temp, original, simulated := runtests(testcase)
	length := len(temp)
	for i := 0; i < length; i++ {
		e := "|" + temp[i].name + "| Result: " + temp[i].result + "| Expected: ~" + temp[i].expected
		fmt.Fprintln(f, e)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	lastln := "\nORIGINAL: " + original + " | USED: " + simulated
	l2, failure := f.WriteString(lastln)
	if failure != nil {
		fmt.Println(failure)
		f.Close()
		return
	}
	fmt.Println(l+l2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// This might be my longest function in this entire repo... don't know yet though
func runtests(testcase string) ([]results, string, string) {
	holder := []results{}
	var original, used string
	if testcase == "CASE0" {
		length := len(gen.Case0)
		for i := 0; i < length; i++ {
			x := gen.Case0[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop1)
			//fmt.Println("USED: ", used, "ORIGINAL: ", original)
			x.Args.Data = Sens1Prop1
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE1" {
		length := len(gen.Case1)
		for i := 0; i < length; i++ {
			x := gen.Case1[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop2)
			x.Args.Data = Sens1Prop2
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE2" {
		length := len(gen.Case2)
		for i := 0; i < length; i++ {
			x := gen.Case2[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop3)
			x.Args.Data = Sens1Prop3
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE3" {
		length := len(gen.Case3)
		for i := 0; i < length; i++ {
			x := gen.Case3[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop4)
			x.Args.Data = Sens1Prop4
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE4" {
		length := len(gen.Case4)
		for i := 0; i < length; i++ {
			x := gen.Case4[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop5)
			x.Args.Data = Sens1Prop5
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE5" {
		length := len(gen.Case5)
		for i := 0; i < length; i++ {
			x := gen.Case5[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop6)
			x.Args.Data = Sens1Prop6
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE6" {
		length := len(gen.Case6)
		for i := 0; i < length; i++ {
			x := gen.Case6[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop7)
			x.Args.Data = Sens1Prop7
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE7" {
		length := len(gen.Case7)
		for i := 0; i < length; i++ {
			x := gen.Case7[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop8)
			x.Args.Data = Sens1Prop8
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE8" {
		length := len(gen.Case8)
		for i := 0; i < length; i++ {
			x := gen.Case8[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop9)
			x.Args.Data = Sens1Prop9
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE9" {
		length := len(gen.Case9)
		for i := 0; i < length; i++ {
			x := gen.Case9[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop10)
			x.Args.Data = Sens1Prop10
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE10" {
		length := len(gen.Case10)
		for i := 0; i < length; i++ {
			x := gen.Case10[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop11)
			x.Args.Data = Sens1Prop11
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE11" {
		length := len(gen.Case11)
		for i := 0; i < length; i++ {
			x := gen.Case11[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop12)
			x.Args.Data = Sens1Prop12
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE12" {
		length := len(gen.Case12)
		for i := 0; i < length; i++ {
			x := gen.Case12[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop13)
			x.Args.Data = Sens1Prop13
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Result)
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE13" {
		length := len(gen.Case13)
		for i := 0; i < length; i++ {
			x := gen.Case13[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop14)
			x.Args.Data = Sens1Prop14
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE14" {
		length := len(gen.Case14)
		for i := 0; i < length; i++ {
			x := gen.Case14[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop15)
			x.Args.Data = Sens1Prop15
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE15" {
		length := len(gen.Case15)
		for i := 0; i < length; i++ {
			x := gen.Case15[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop16)
			x.Args.Data = Sens1Prop16
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE16" {
		length := len(gen.Case16)
		for i := 0; i < length; i++ {
			x := gen.Case16[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop17)
			x.Args.Data = Sens1Prop17
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE17" {
		length := len(gen.Case17)
		for i := 0; i < length; i++ {
			x := gen.Case17[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop18)
			x.Args.Data = Sens1Prop18
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE18" {
		length := len(gen.Case18)
		for i := 0; i < length; i++ {
			x := gen.Case18[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop19)
			x.Args.Data = Sens1Prop19
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE19" {
		length := len(gen.Case19)
		for i := 0; i < length; i++ {
			x := gen.Case19[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop20)
			x.Args.Data = Sens1Prop20
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	} else if testcase == "CASE20" {
		length := len(gen.Case20)
		for i := 0; i < length; i++ {
			x := gen.Case20[i]
			original, used = fmt.Sprintf("%v", x.Args.Data), fmt.Sprintf("%v", Sens1Prop21)
			x.Args.Data = Sens1Prop21
			exp := fmt.Sprintf("%f", x.Result)
			res := fmt.Sprintf("%f", x.Method(x.Args.Data))
			holder = append(holder, results{name: x.Name, expected: exp, result: res})
		}
	}
	//return holder
	return holder, original, used
}
