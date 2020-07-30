package tcpserver

import (
	"fmt"
	"os"
	"strconv"
)

// DetectSensor : Find out which sensor is being used
func DetectSensor(buf []byte) string {
	var sens string
	length := len(buf)
	for i := 0; i < length; i++ {
		if buf[i] == 44 && buf[i+1] == 103 {
			sens = string(buf[i+1 : i+11])
			break
		}
	}
	return sens
}

// HandleData : Takes in data values and returns an array.
func HandleData(buf []byte) { // this works if there are multiple properties or not
	sensor := DetectSensor(buf)
	idx, prop := 0, 0
	length := len(buf)
	for i := 0; i < length; i++ {
		if buf[i] == 44 {
			temp := string(buf[idx:i])
			val, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				fmt.Println("Error Reading:", err.Error())
				fmt.Println("=====SIMULATION ERROR=====")
				os.Exit(1)
			} else {
				AppendValue(sensor, val, prop)
			}
			idx = i + 1
			prop++
		}
	}
}

// AppendValue : Places correct value in the correct slice
func AppendValue(sensor string, val float64, prop int) {
	if sensor == "gr3sensor1" {
		if prop == 0 {
			Sens1Prop1 = append(Sens1Prop1, val)
			fmt.Println("Sensor 1 Property 1:  ", Sens1Prop1, " case0")
		} else if prop == 1 {
			Sens1Prop2 = append(Sens1Prop2, val)
			fmt.Println("Sensor 1 Property 2:  ", Sens1Prop2, " case1")
		} else if prop == 2 {
			Sens1Prop3 = append(Sens1Prop3, val)
			fmt.Println("Sensor 1 Property 3:  ", Sens1Prop3, " case2")
		} else if prop == 3 {
			Sens1Prop4 = append(Sens1Prop4, val)
			fmt.Println("Sensor 1 Property 4:  ", Sens1Prop4, " case3")
		} else if prop == 4 {
			Sens1Prop5 = append(Sens1Prop5, val)
			fmt.Println("Sensor 1 Property 5:  ", Sens1Prop5, " case4")
		} else if prop == 5 {
			Sens1Prop6 = append(Sens1Prop6, val)
			fmt.Println("Sensor 1 Property 6:  ", Sens1Prop6, " case5")
		} else if prop == 6 {
			Sens1Prop7 = append(Sens1Prop7, val)
			fmt.Println("Sensor 1 Property 7:  ", Sens1Prop7, " case6")
		} else if prop == 7 {
			Sens1Prop8 = append(Sens1Prop8, val)
			fmt.Println("Sensor 1 Property 8:  ", Sens1Prop8, " case7")
		} else if prop == 8 {
			Sens1Prop9 = append(Sens1Prop9, val)
			fmt.Println("Sensor 1 Property 9:  ", Sens1Prop9, " case8")
		} else if prop == 9 {
			Sens1Prop10 = append(Sens1Prop10, val)
			fmt.Println("Sensor 1 Property 10: ", Sens1Prop10, " case9")
		} else if prop == 10 {
			Sens1Prop11 = append(Sens1Prop11, val)
			fmt.Println("Sensor 1 Property 11: ", Sens1Prop11, " case10")
		} else if prop == 11 {
			Sens1Prop12 = append(Sens1Prop12, val)
			fmt.Println("Sensor 1 Property 12: ", Sens1Prop12, " case11")
		} else if prop == 12 {
			Sens1Prop13 = append(Sens1Prop13, val)
			fmt.Println("Sensor 1 Property 13: ", Sens1Prop13, " case12")
		} else if prop == 13 {
			Sens1Prop14 = append(Sens1Prop14, val)
			fmt.Println("Sensor 1 Property 14: ", Sens1Prop14, " case13")
		} else if prop == 14 {
			Sens1Prop15 = append(Sens1Prop15, val)
			fmt.Println("Sensor 1 Property 15: ", Sens1Prop15, " case14")
		} else if prop == 15 {
			Sens1Prop16 = append(Sens1Prop16, val)
			fmt.Println("Sensor 1 Property 16: ", Sens1Prop16, " case15")
		} else if prop == 16 {
			Sens1Prop17 = append(Sens1Prop17, val)
			fmt.Println("Sensor 1 Property 17: ", Sens1Prop17, " case16")
		} else if prop == 17 {
			Sens1Prop18 = append(Sens1Prop18, val)
			fmt.Println("Sensor 1 Property 18: ", Sens1Prop18, " case17")
		} else if prop == 18 {
			Sens1Prop19 = append(Sens1Prop19, val)
			fmt.Println("Sensor 1 Property 19: ", Sens1Prop19, " case18")
		} else if prop == 19 {
			Sens1Prop20 = append(Sens1Prop20, val)
			fmt.Println("Sensor 1 Property 20: ", Sens1Prop20, " case19")
		} else if prop == 20 {
			Sens1Prop21 = append(Sens1Prop21, val)
			fmt.Println("Sensor 1 Property 21: ", Sens1Prop21, " case20")
		}
	}
}
