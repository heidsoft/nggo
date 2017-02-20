package go_design_pattern

//定义目标(发布者),用于被观察者观察或订阅
type Publisher interface {
	Publish(value interface{})
}

//定义观察者
type Observer interface {
	Notify(value interface{})
}

//定义函数类型
type ObserverFunc func(value interface{})

func (fn ObserverFunc) Notify(value interface{}){
	fn(value)
}

//定义可用的观察者
type Observable []Observer

func (observers Observable) AddObserver(a Observer){
	observers = append(observers,a)
}