package main

import (
	"fmt"
	"time"
)

func goRoutineA(a <-chan int) {
	val := <-a
	fmt.Println("goRoutineA received the data", val)
}
func main() {
	ch := make(chan int)
	go goRoutineA(ch)
	time.Sleep(time.Second * 1)

	ch <- 11

	time.Sleep(time.Second * 1)

}
