package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:3040")
	conn, _ := listener.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Println(string(message))
		time.Sleep(1 * time.Second)
	}
}
