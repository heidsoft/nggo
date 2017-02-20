package main

import "fmt"
import "github.com/dspinhirne/netaddr-go"

func main() {
	net,_ := netaddr.ParseIPv4Net("192.168.1.0/24")
	fmt.Println(net.String())
	fmt.Println(net.Netmask())
	fmt.Println(net.Network())


}