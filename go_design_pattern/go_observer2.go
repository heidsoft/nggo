package main

import "fmt"

type Publisher interface {
	Publish(value interface{})
}

type Observer interface {
	Notify(value interface{})
}

type ObserverFunc func(value interface{})

//实现Notify方法
func (fn ObserverFunc) Notify(value interface{}){
	fn(value)
}

type Observable []Observer

func (observers *Observable) AddObserver(a Observer){
	*observers = append(*observers, a)
}

func (observers Observable) Publish(value interface{}){
	for _, obs := range observers {
		obs.Notify(value)
	}
}

type Field struct {
	Value int64
	Observable
}

func (f *Field) Set(v int64){
	f.Value = v
	//f.Publish(v)
	f.Observable.Publish(v)
}

func Listen(value interface{}){
	fmt.Printf("new value 1: %v\n", value)
}

func Listen2(value interface{}){
	fmt.Printf("new value 2: %v\n", value)
}

func main() {

	fmt.Println(ObserverFunc(Listen))
	//创建目标
	v := &Field{}
	//加入观察器
	v.AddObserver(ObserverFunc(Listen))
	v.AddObserver(ObserverFunc(Listen2))
	v.Set(105)

	fmt.Println("Hello, playground")
}