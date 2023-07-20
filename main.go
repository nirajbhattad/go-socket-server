// server.go
package main

import (
	"fmt"
	"log"
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

	// Handle incoming client connections
	log.Println("Server starting to listen on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
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
			fmt.Println("Error reading message from shadow:", err.Error())
			break
		}

		// Process client request
		request := string(buffer[:bytesRead])
		log.Println("Received request from shadow:", request)

		// Send response back to the client
		response := "Message Received"
		conn.Write([]byte(response))

		// Clear the buffer
		for i := 0; i < len(buffer); i++ {
			buffer[i] = 0
		}
	}
	conn.Close()
}
