package main

import (
	"fmt"
	"time"
)

func main() {
	ticker1 := time.NewTicker(1*time.Second)
	for{
		select {
		case <-ticker1.C:
			go func() {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			}()
		}
	}
}

func TimeBuild(strTime string) time.Time {
	tm, _ := time.Parse("2006-01-02 15:04:05", strTime)
	return tm
}

