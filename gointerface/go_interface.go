package main

import (
	"fmt"
	"io"
)

type TestInterface interface {
	Test1(test string)
//	Test2(test string) string
}

type TestStruct struct {
	Name string
}

func (this *TestStruct)  Test1() {
	fmt.Println("test1....")
}

//func (this *TestStruct)  Test2() string{
//	fmt.Println("test2....")
//	return "test2"
//}

type TestCreateFactory interface {
	Create()  TestInterface
}

type TestStructFactory struct {

}

func (t *TestStructFactory)  Create() TestInterface{
	return  nil
}


type Shape interface {
	Draw(io.Writer) error
}

type Point struct {

}

type Circle struct {
	Location Point
	Radius   float64
}

func (c *Circle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<circle cx="%f" cy="%f" r="%f"/>`, c.Location.X, c.Location.Y, c.Radius)
	return err
}

func main()  {
	test := TestStruct{
		Name:"test",
	}

	test.Test1()
}