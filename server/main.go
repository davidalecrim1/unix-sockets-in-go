package main

import (
	"fmt"
	"net"
	"os"
)

type Server interface {
	Start() error
	handleConnection(conn net.Conn)
}

type UnixServer struct {
	socketPath string
}

func NewUnixServer(socketPath string) *UnixServer {
	return &UnixServer{
		socketPath: socketPath,
	}
}

func (s *UnixServer) Start() error {

	// Ensure the socket does not already exist
	if err := os.RemoveAll(s.socketPath); err != nil {
		return fmt.Errorf("error removing existing socket: %w", err)
	}

	listener, err := net.Listen("unix", s.socketPath)

	if err != nil {
		return fmt.Errorf("error creating Unix Socket: %w", err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on: ", s.socketPath)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate goroutine
		go s.handleConnection(conn)
	}
}

func (s *UnixServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// the buffer supports 1024 characters given each character is one byte
	buf := make([]byte, 1024)

	// n is the number of bytes that were actually read into buf from the connection. 
	// This value is returned by the Read method.
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// This creates a slice of buf from index 0 to n, 
	// effectively trimming the buffer to only include the actual data read.
	message := string(buf[:n])
	fmt.Println("Received message: ", message)
	response := "Message received! The message was: " + message 
	
	if _, err = conn.Write([]byte(response)); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func main(){
	socketPath := os.Getenv("GO_SOCKET_PATH")

	if socketPath == "" {
		fmt.Println("GO_SOCKET_PATH environment variable is not set. Using default socket path.")
		socketPath = "/tmp/go/unixsocket"
	}

	server := NewUnixServer(socketPath)

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
	}

}