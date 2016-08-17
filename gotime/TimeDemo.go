package main

import (
	"fmt"
	"time"
)

func main() {
	//时间转换
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	tm, _ := time.Parse("2006-01-02 15:04:05", "2016-06-02 15:16:20")
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
	fmt.Println(tm.Unix())

	tm2 := TimeBuild("2016-06-02 15:16:20")
	fmt.Println(tm2.Unix())
	fmt.Println(tm2.Format("2006-01-02 15:04:05"))

	fmt.Println("aaaaaaaa")
	fmt.Println(time.Parse("2006-01-02 15:04:05",time.))

}

func TimeBuild(strTime string) time.Time {
	tm, _ := time.Parse("2006-01-02 15:04:05", strTime)
	return tm
}

