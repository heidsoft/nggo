package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main() {
	AppPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(AppPath)
	os.Chdir(AppPath)
}
