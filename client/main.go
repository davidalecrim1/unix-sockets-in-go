package main

import (
	"fmt"
	"net"
	"os"
)

type Client interface {
	Connect() error
	sendMessage(message string) error
	readMessage() (string, error)
}

type UnixClient struct {
	socketPath string
	conn net.Conn
}

func NewUnixClient(socketPath string) *UnixClient {
	return &UnixClient{
		socketPath: socketPath,
	}
}

func (c *UnixClient) Connect() error {
	conn, err := net.Dial("unix", c.socketPath)
	
	if err != nil {
		return fmt.Errorf("error connecting to Unix socket: %w", err)
	}

	c.conn = conn
	return nil
}

func (c *UnixClient) sendMessage(message string) error {
	_, err := c.conn.Write([]byte(message))

	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	return nil
}

func (c *UnixClient) readMessage() (string, error){
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)

	if err != nil {
		return "", fmt.Errorf("error reading message: %w", err)
	}

	return string(buffer[:n]), nil
}

func main(){
	socketPath := os.Getenv("GO_SOCKET_PATH")

	if socketPath == "" {
		fmt.Println("GO_SOCKET_PATH environment variable is not set. Using default socket path.")
		socketPath = "/tmp/go/unixsocket"
	}

	client := NewUnixClient(socketPath)

	if err := client.Connect(); err != nil {
		fmt.Println("Client connection error: ", err)
		return
	}

	defer client.conn.Close()
	fmt.Println("Connected to Unix socket")

	message := "Hello, Server. This is the client 01" //36 characters
	
	if err := client.sendMessage(message); err != nil {
		fmt.Println("Error sending message: ", err)
	}

	reply, err := client.readMessage()

	if err != nil {
		fmt.Println("Failed to read the message: ", err)
		return
	}
	
	fmt.Println("Received a reply message: ", reply)
}