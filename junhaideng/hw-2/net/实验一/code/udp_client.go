package main

import (
	"fmt"
	"net"
)

func main(){
	udpAddr, err := net.ResolveUDPAddr("udp",":8888")
	if err != nil{
		fmt.Println("Err resolve UDP address: ", err)
		return 
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil{
		fmt.Println("Dial UDP error: ", err)
		return 
	}

	for {
		conn.Write([]byte("Hello server"))
		var buff [512]byte 
		n, err := conn.Read(buff[0:])
		if err != nil{
			fmt.Println("ERR: ", err)
			break
		}
		fmt.Println("Read from server: ", string(buff[:n]))
	}

}