package main

import (
	"fmt"
	"time"
)

/**
https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
*/
func main() {
	fmt.Printf("main() started")
	c := make(chan string, 2)
	go greet(c)

	c <- "Jake test"

	time.Sleep(time.Second * 2)

	fmt.Printf("main() stopted")
}

func greet(c chan string) {
	fmt.Println("hello " + <-c + "|")
}
