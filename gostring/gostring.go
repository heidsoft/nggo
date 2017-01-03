package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(strings.ContainsAny("111", "aaabbbb"))
	fmt.Println("控制结果1", strings.Contains("select", "manageserver"))
	fmt.Println("控制结果2", strings.Contains("manage", "manageserver"))
	fmt.Println("控制结果3", strings.Contains("system", "manageserver"))
	fmt.Println("控制结果3", strings.Contains("manageserver", "manage"))
	fmt.Println("控制结果3", strings.Contains("manageserver", "system"))
	fmt.Println("控制结果3", strings.Contains("manageserver", "select"))

	replaceString := strings.Replace("dcname.", ".", "__", -1)
	joinString := strings.Join([]string{"aaa"}, "bbb")
	fmt.Println("replaceString", replaceString)
	fmt.Println("joinString", joinString)
}