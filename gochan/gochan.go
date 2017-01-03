package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	c := make(chan int)

	// 使写操作在另一个goroutine中执行。
	go func() {
		c <- 42
	}()
	val := <-c
	fmt.Println(val)
}
