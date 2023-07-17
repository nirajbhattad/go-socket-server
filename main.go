// server.go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Start the server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on :8080")

	// Handle incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		// Read client request
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		// Process client request
		request := string(buffer[:bytesRead])
		fmt.Println("Received request:", request)

		// Send response back to the client
		response := "Hello from server"
		conn.Write([]byte(response))

		// Clear the buffer
		for i := 0; i < len(buffer); i++ {
			buffer[i] = 0
		}
	}
	conn.Close()
}
