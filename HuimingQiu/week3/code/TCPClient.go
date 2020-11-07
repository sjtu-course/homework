package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	msg := flag.String("msg", "Hello", "The message to send")
	flag.Parse()
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte(*msg))
	if err != nil {
		log.Fatal()
	}

	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	fmt.Printf("Receive: ", string(buffer[:n]))
}
