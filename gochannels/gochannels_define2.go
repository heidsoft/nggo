package main

import "fmt"

//https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
func main() {
	c := make(chan int)

	fmt.Printf("Type of `c` is %T\n", c)
	fmt.Printf("Value of `c  is %v\n", c)
}
