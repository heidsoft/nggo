package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	message := make(chan string)
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			message <- fmt.Sprintf("message %d", i)

		}
		//close(message)
	}()


	//time.Sleep(time.Second*6)
	//for msg := range message {
	//	fmt.Println(msg)
	//}

	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
}
