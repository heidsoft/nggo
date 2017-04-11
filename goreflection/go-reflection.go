package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)



//go 反射测试
func main() {

	tst := "string"
	tst2 := 10
	tst3 := 1.2



	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))
	fmt.Println(reflect.ValueOf(tst))
	fmt.Println(reflect.ValueOf(tst2))
	fmt.Println(reflect.ValueOf(tst3))

	type MetaData struct {
		Key   string `json:"key,omitempty"`
		Value map[interface{}]interface{} `json:"value,omitempty"`
	}

	type Cpu struct {
		name string
	}

	value := make(map[interface{}]interface{})
	value["cpu"] = "11"
	value["mem"] = "12"
	//value = append(value,Cpu{name:"test1"})

	testMetaData := MetaData{
		Key:"test",
		Value:value,
	}

	fmt.Println(reflect.TypeOf(testMetaData.Value).Name())
	jsonBytes, _ := json.Marshal(testMetaData.Value)
	mycpu := new(Cpu)
	json.Unmarshal(jsonBytes,mycpu)
	fmt.Println("111111")
	fmt.Println(mycpu.name)


}