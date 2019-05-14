package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 2)

	ch <- 100
	ch2 <- 200

	for i := 0; i < 2; i++ {
		select {
		case a := <-ch:
			fmt.Printf("ch send msg %d\n", a)
		case b := <-ch2:
			fmt.Printf("ch2 send msg %d\n", b)
		default:
			fmt.Printf("no  msg \n")
		}
	}

	time.Sleep(time.Second * 5)
}
