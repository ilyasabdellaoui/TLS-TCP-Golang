package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	clientCert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Printf("Ops! failed to load client certificates: %v\n", err)
		return
	}

	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		fmt.Printf("Ops! failed to load CA certificate: %v\n", err)
		return
	}

	serverCertPool := x509.NewCertPool()
	serverCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            serverCertPool,
		InsecureSkipVerify: false,
	}

	conn, err := tls.Dial("tcp", "localhost:443", tlsConfig)
	if err != nil {
		fmt.Printf("Ops! failed to connect to server: %v", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to secure server!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Message: ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ops! failed to read input: %v", err)
			break
		}

		msg = msg[:len(msg)-1]

		if msg == "exit" {
			break
		}

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("Ops! failed to send message: %v", err)
			break
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Ops! an error has occurred while reading data: %v", err)
			break
		}

		fmt.Printf("Message from server: %s\n", string(buffer[:n]))
	}
}
