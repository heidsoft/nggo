package main

import (
	"os"
	"fmt"
	"path/filepath"
)


func main() {
	fmt.Println("args1 ==> ",os.Args[0])
	appPath,_:= filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("AppPath ==> ",appPath)

	err := os.Chdir("/Users/heidsoft")
	if err != nil{
		fmt.Println("error ==> ",err.Error())
	}

	file, err := os.Create("test.txt")
	if err == nil {
		fmt.Println("name ==> ",file.Name())
	}

	// 0000 0000 0000 0000
	fmt.Println(1<<26/1024/1024)
	fmt.Println(1<<1)
	fmt.Println(1>>2)
}
