package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(connection net.Conn) {
	buffer := make([]byte, 256)
	n, err := connection.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	msg := string(buffer[:n])
	fmt.Println("Receive: ", string(msg))
	n, _ = connection.Write([]byte("ACK: " + msg))
}

func main() {
	TCPAddr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", TCPAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		connection, _ := listener.Accept()

		go handleConnection(connection)
	}
}
