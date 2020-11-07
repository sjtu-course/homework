package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "localhost:8000" // 远程服务器的ip地址和端口

	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	conn, err := net.DialUDP("udp", nil, udpAddr)

	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}

	conn.Write([]byte("Hello World!"))

	fmt.Println("send msg")

	var msg [20]byte
	n, err := conn.Read(msg[0:])

	fmt.Println("msg is:", string(msg[:n]))
}
