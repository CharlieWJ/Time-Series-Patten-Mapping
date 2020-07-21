package tcpserver

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"

	gen "../generatedInGo" //golint sometimes tells me there is an error but i don't recieve one when running, not sure why
)

//ASK ABOUT PACKAGING ISSUES
const (
	connHost = "localhost"
	connPort = "7100"
	connType = "tcp"
)

var datalength = 0

var sens1Prop1 = []float64{}
var sens1Prop2 = []float64{}
var sens1Prop3 = []float64{}
var sens1Prop4 = []float64{}
var sens1Prop5 = []float64{}
var sens1Prop6 = []float64{}

var datasens2 = []float64{}
var datasens3 = []float64{}
var datasens4 = []float64{}
var datasens5 = []float64{}
var datasens6 = []float64{}

// Float64frombytes : Takes a byte array and converts it to a float64
func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// Float64bytes : converts float64 to byte array/slice
func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// HandleRequest : Handles incoming requests.
func HandleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 64) //make([]byte, 1024)
	// Read the incoming connection into the buffer.
	for {
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error Reading:", err.Error())
			fmt.Println("=====SIMULATION HAS CONCLUDED=====")
			os.Exit(1)
		}
		_ = reqLen // Unused variable
		fmt.Println(string(buf))
		HandleData(buf)
		fmt.Println("")
		if datalength+10 == len(sens1Prop1) {
			fmt.Println("=====FUNCTION TESTING=====")
			fmt.Println("Max_range_decreasing(): ", gen.Max_range_decreasing(sens1Prop1))
			fmt.Println("Max_range_decreasing_sequence(): ", gen.Max_range_decreasing_sequence(sens1Prop1))
			fmt.Println("Max_range_strictly_decreasing_sequence(): ", gen.Max_range_strictly_decreasing_sequence(sens1Prop1))
			fmt.Println("Max_range_increasing(): ", gen.Max_range_increasing(sens1Prop2))
			fmt.Println("Max_range_increasing_sequence(): ", gen.Max_range_increasing_sequence(sens1Prop2))
			fmt.Println("Max_range_strictly_increasing_sequence(): ", gen.Max_range_strictly_increasing_sequence(sens1Prop2))
			datalength = len(sens1Prop1)
		}
		//conn.Write([]byte("Message received."))
	}
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

// AppendValue : Places correct value in the correct slice
func AppendValue(sensor string, val float64, prop int) {
	if sensor == "gr3sensor1" {
		if prop == 0 {
			sens1Prop1 = append(sens1Prop1, val)
			fmt.Println("Sensor 1 Property 1: ", sens1Prop1)
		} else if prop == 1 {
			sens1Prop2 = append(sens1Prop2, val)
			fmt.Println("Sensor 1 Property 2: ", sens1Prop2)
		} else if prop == 2 {
			sens1Prop3 = append(sens1Prop3, val)
			fmt.Println("Sensor 1 Property 3: ", sens1Prop3)
		} else if prop == 3 {
			sens1Prop3 = append(sens1Prop4, val)
			fmt.Println("Sensor 1 Property 4: ", sens1Prop4)
		} else if prop == 4 {
			sens1Prop3 = append(sens1Prop5, val)
			fmt.Println("Sensor 1 Property 5: ", sens1Prop5)
		} else {
			sens1Prop6 = append(sens1Prop6, val)
			fmt.Println("Sensor 1 Property 6: ", sens1Prop6)
		}
	} else if sensor == "gr3sensor2" {
		datasens2 = append(datasens2, val)
		fmt.Println("Sensor 2: ", datasens2)
	} else if sensor == "gr3sensor3" {
		datasens3 = append(datasens3, val)
		fmt.Println(datasens3)
	} else if sensor == "gr3sensor4" {
		datasens4 = append(datasens4, val)
		fmt.Println(datasens4)
	} else if sensor == "gr3sensor5" {
		datasens5 = append(datasens5, val)
		fmt.Println(datasens5)
	} else {
		datasens6 = append(datasens6, val)
		fmt.Println(datasens6)
	}
}

// RunServer : Calls 'HandleRequest' and keeps server up and running
func RunServer() {
	// Listen for incoming connections.
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		fmt.Println("=====SIMULATION HAS CONCLUDED=====")
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer listener.Close()
	fmt.Println("Listening on " + connHost + ":" + connPort)
	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			fmt.Println("=====SIMULATION HAS CONCLUDED=====")
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go HandleRequest(conn)
	}
}
