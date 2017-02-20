package main

import (
	"os"
	"fmt"
	"os/signal"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,os.Kill)
	done := make(chan bool, 1)

	//SIGINT ctr+c
	go func() {
		fmt.Println("信号处理")
		sig := <-c
		fmt.Println(sig)
		done <- true
	}()

	<-done
	fmt.Println("exiting")
}
