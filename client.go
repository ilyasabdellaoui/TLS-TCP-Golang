package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", "localhost:443", tlsConfig)
	if err != nil {
		fmt.Printf("Ops! failed to connect to server: %v", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to secure server!")
	reader := bufio.NewReader(os.Stdin)

	for {
		var msg string
		fmt.Print("Message: ")
		msg, err := reader.ReadString('\n')

		msg = msg[:len(msg)-1]

		if msg == "Bye!" {
			conn.Write([]byte(msg))
			break
		}

		conn.Write([]byte(msg))

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Ops! an error has occured while reading data: %v", err)
			break
		}

		fmt.Printf("Message from server: %s\n", string(buffer[:n]))
	}
}
