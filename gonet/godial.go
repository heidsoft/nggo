package main

import (
	"net"
	"fmt"
	"os"
)


func main()  {
	//创建TCP4地址
	tcpAddr,err := net.ResolveTCPAddr("tcp4","localhost:80")
	if err != nil {
		fmt.Fprintf(os.Stderr,"error1==> %s",err.Error())
	}
	fmt.Println(tcpAddr.String())
	fmt.Println(tcpAddr.Port)
	fmt.Println(tcpAddr.IP)
	fmt.Println(tcpAddr.Network())

	tcpConn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr,"error2==> %s",err.Error())
	}
	numByte,err := tcpConn.Write([]byte("abc1234567890"))
	if err != nil {
		fmt.Fprintf(os.Stderr,"error3==> %s",err.Error())
	}
	fmt.Fprintf(os.Stdout,"out==> %d",numByte)



}
