package main

import (
    "os"
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("udp", "127.0.0.1:8080")
    defer conn.Close()
    if err != nil {
        os.Exit(1)  
    }

    conn.Write([]byte("Hello world!"))

    fmt.Println("send msg")

    var msg [50]byte
		n, err := conn.Read(msg[0:])
		if err != nil{
			fmt.Println("ERR: ", err)
		}

    fmt.Println("msg is", string(msg[0:n]))
}