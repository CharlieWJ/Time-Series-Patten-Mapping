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

const (
	// ConnHost : Connection host operated on
	ConnHost = "localhost"
	// ConnPort : Port where connection is established
	ConnPort = "7100"
	// ConnType : Connection Type
	ConnType = "tcp"
)

// DataSens1 : Slice that corresponds the sensor sending data
var DataSens1 = []float64{}
var d1len = len(DataSens1)

// Sens1Prop1 : Sensor 1 Property 1
var Sens1Prop1 = []float64{}
var s1p1len = len(Sens1Prop1)

// Sens1Prop2 : Sensor 1 Property 2
var Sens1Prop2 = []float64{}
var s1p2len = len(Sens1Prop2)

// Sens1Prop3 : Sensor 1 Property 3
var Sens1Prop3 = []float64{}
var s1p3len = len(Sens1Prop3)

// DataSens2 : Slice that corresponds the sensor sending data
var DataSens2 = []float64{}

// DataSens3 : Slice that corresponds the sensor sending data
var DataSens3 = []float64{}

// DataSens4 : Slice that corresponds the sensor sending data
var DataSens4 = []float64{}

// DataSens5 : Slice that corresponds the sensor sending data
var DataSens5 = []float64{}

// DataSens6 : Slice that corresponds the sensor sending data
var DataSens6 = []float64{}

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
		//fmt.Println(string(buf))
		//fmt.Println(buf)
		var s string
		var sensor string
		for i := 0; i < len(buf); i++ {
			if buf[i] == 44 && buf[i+1] == 103 { // Detects comma and 'g'
				temp := buf[0:i]
				s = string(temp)
				sensor = string(buf[i+1 : i+11]) // determines which sensor
				val, err2 := strconv.ParseFloat(s, 64)
				if err2 != nil {
					fmt.Println("Error Reading:", err.Error())
					fmt.Println("=====SIMULATION HAS CONCLUDED=====")
					os.Exit(1)
				}
				if sensor == "gr3sensor1" {
					DataSens1 = append(DataSens1, val)
					fmt.Print(DataSens1)
				} else if sensor == "gr3sensor2" {
					DataSens2 = append(DataSens2, val)
					fmt.Print(DataSens2)
				} else if sensor == "gr3sensor3" {
					DataSens3 = append(DataSens3, val)
					fmt.Print(DataSens3)
				} else if sensor == "gr3sensor4" {
					DataSens4 = append(DataSens4, val)
					fmt.Print(DataSens4)
				} else if sensor == "gr3sensor5" {
					DataSens5 = append(DataSens5, val)
					fmt.Print(DataSens5)
				} else {
					DataSens6 = append(DataSens6, val)
					fmt.Print(DataSens6)
				}
			}
		}
		/*val, err2 := strconv.ParseFloat(s, 64)
		if err2 != nil {
			fmt.Println("Error Reading:", err.Error())
			fmt.Println("=====SIMULATION HAS CONCLUDED=====")
			os.Exit(1)
		}
		// Below is sensor detection
		if sensor == "gr3sensor1" {
			DataSens1 = append(DataSens1, val)
			fmt.Print(DataSens1)
		} else if sensor == "gr3sensor2" {
			DataSens2 = append(DataSens2, val)
			fmt.Print(DataSens2)
		} else if sensor == "gr3sensor3" {
			DataSens3 = append(DataSens3, val)
			fmt.Print(DataSens3)
		} else if sensor == "gr3sensor4" {
			DataSens4 = append(DataSens4, val)
			fmt.Print(DataSens4)
		} else if sensor == "gr3sensor5" {
			DataSens5 = append(DataSens5, val)
			fmt.Print(DataSens5)
		} else {
			DataSens6 = append(DataSens6, val)
			fmt.Print(DataSens6)
		}
		*/
		fmt.Println(", ", sensor)
		fmt.Println("")
		if d1len+18 == len(DataSens1) {
			fmt.Println("Max_range_decreasing_sequence(", DataSens1, "): ", gen.Max_range_decreasing_sequence(DataSens1))
			d1len = len(DataSens1)
		}
		//conn.Write([]byte("Message received."))
	}
}

// HandleData : Takes in data values and returns an array.
func HandleData(data []float64, val float64) []float64 {
	return (append(data, val))
}

// RunServer : Calls 'HandleRequest' and keeps server up and running
func RunServer() {
	// Listen for incoming connections.
	listener, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		fmt.Println("=====SIMULATION HAS CONCLUDED=====")
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer listener.Close()
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
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
