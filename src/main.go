package main

import (
	"fmt"
	"math"

	tcp "./TCPServer"
	//generated "./generatedInGo"
)

const (
	ConnHost = "localhost"
	ConnPort = "7100"
	ConnType = "tcp"
)

func add(x float64, y float64) float64 {
	return (x + y)
}
func diff(x float64, y float64) float64 { return (math.Abs(x - y)) }

/*
// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	for {
		reqLen, err := conn.Read(buf)
		_ = reqLen // Unused variable
		if err != nil {
			fmt.Println("Error Reading:", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(buf))
	}
	// Close the connection when you're done with it.
	// conn.Close()
}
*/
func main() {
	tcp.RunServer()
	//fmt.Println(generated.Max_range_decreasing_sequence([]float64{5, 3, 2, 5}))
	fmt.Println("=====Simulation has Concluded=====")
}
