package main

import (
	"net"
	"fmt"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println("error:",err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		b := make([]byte,0)
		conn.Read(b)

	}
}
