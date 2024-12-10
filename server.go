package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Printf("Ops! failed to load certificates: %v\n", err)
		return
	}

	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		fmt.Printf("Ops! failed to load CA certificate: %v\n", err)
		return
	}

	clientCAs := x509.NewCertPool()
	clientCAs.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCAs,
	}

	listener, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		fmt.Printf("Ops! failed to start server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Secure server is running on port 443! Waiting for clients...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ops! an error has occurred while accepting connection: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from secure client: ", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Ops! an error has occurred while reading data: ", err)
			}
			break
		}

		clMsg := string(buffer[:n])
		clMsg = strings.TrimSpace(clMsg)
		fmt.Printf("Message from client %s: %s\n", conn.RemoteAddr(), clMsg)

		if clMsg == "who are you?" {
			_, err := conn.Write([]byte("I am secure server!"))
			if err != nil {
				fmt.Println("Ops! failed to send response: ", err)
				return
			}
		} else if clMsg == "exit" {
			break
		} else {
			_, err := conn.Write([]byte("I don't understand!"))
			if err != nil {
				fmt.Println("Ops! failed to send response: ", err)
				return
			}
		}
	}
	fmt.Println("Client disconnected: ", conn.RemoteAddr())
}
