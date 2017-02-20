//ip模块处理测试

package main

import (
	"net"
	"fmt"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

//解析cidr
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	fmt.Println("mask:",ipnet.Mask)
	prefixSize, _ := ipnet.Mask.Size()
	fmt.Println(prefixSize)
	if err != nil {
		return nil, err
	}


	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func main()  {
	ips,err := Hosts("10.0.2.24/16")
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("ip的长度是",len(ips))
	//for index := range ips{
	//	fmt.Println(ips[index])
	//}
}