package main

import "fmt"
import "reflect"

func main() {
	abc := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	keys := reflect.ValueOf(abc).MapKeys()

	fmt.Println(keys) // [a b c]
}
