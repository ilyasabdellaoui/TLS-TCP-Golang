package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Printf("Ops! failed to load certificates: %v", err)
		return
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		fmt.Printf("Ops! failed to start server: %v", err)
		return
	}
	defer listener.Close()

	fmt.Println("Secure server is running on port 443! Waiting for clients...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ops! an error has occured while accepting connection: ", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from secure client: ", conn.RemoteAddr())

	// if the server got the msg who are you?, it responds with its name
	// if the server got the msg exit, it sends bye to the client and closes the connx
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Ops! an error has occured while reading data: ", err)
			}
			break
		}

		clMsg := string(buffer[:n])
		clMsg = strings.TrimSpace(clMsg)
		fmt.Printf("Message from client %s: %s\n", conn.RemoteAddr(), clMsg)

		if clMsg == "who are you?" {
			conn.Write([]byte("I am secure server!"))
		} else {
			conn.Write([]byte("I don't understand!"))
		}
	}
	fmt.Println("Client disconnected: ", conn.RemoteAddr())
}
