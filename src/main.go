package main

import (
	"fmt"
	"math"
	"net"
	"os"
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
		// Send a response back to person contacting us.
		//conn.Write([]byte("Message received."))

		fmt.Println(string(buf))
	}
	// Close the connection when you're done with it.
	// conn.Close()
}

func main() {
	// Listen for incoming connections.
	listener, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer listener.Close()
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
	//fmt.Println("Gets Here")
	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
		//fmt.Printf("conn: %#p \n", conn)
		//temp := conn
		//fmt.Printf("%v\n----------\n", temp)
		//fmt.Println(temp)
	}
}
