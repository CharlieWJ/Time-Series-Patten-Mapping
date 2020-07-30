package tcpserver

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"os"
	//golint sometimes tells me there is an error but i don't recieve one when running, not sure why
)

//ASK ABOUT PACKAGING ISSUES
const (
	connHost = "localhost"
	connPort = "7100"
	connType = "tcp"
)

var datalength = 0

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
		if datalength+16 == len(Sens1Prop1) {
			WriteToFile("CASE0")
			WriteToFile("CASE1")
			WriteToFile("CASE2")
			WriteToFile("CASE3")
			WriteToFile("CASE4")
			WriteToFile("CASE5")
			WriteToFile("CASE6")
			WriteToFile("CASE7")
			WriteToFile("CASE8")
			WriteToFile("CASE9")
			WriteToFile("CASE10")
			WriteToFile("CASE11")
			WriteToFile("CASE12")
			WriteToFile("CASE13")
			WriteToFile("CASE14")
			WriteToFile("CASE15")
			WriteToFile("CASE16")
			WriteToFile("CASE17")
			WriteToFile("CASE18")
			WriteToFile("CASE19")
			WriteToFile("CASE20")
			datalength += 16
		}
		//conn.Write([]byte("Message received."))
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
