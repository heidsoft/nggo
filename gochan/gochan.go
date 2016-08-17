package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	ci := make(chan int)
	//cs :=make(chan string)
	//cf :=make(chan interface{})

	ci <- 1
	fmt.Sprintf("%", <-ci)

	//cs <- "test"
	//fmt.Println(<-cs)
	//
	//cf <- user{
	//	name:"liubin",
	//	age:12,
	//}
	//
	//fmt.Println(<-cf)

}
