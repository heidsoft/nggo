package main
import (
	"fmt"
	"reflect" // 这里引入reflect模块
)

type User struct {
	Name   string "user name" //这引号里面的就是tag
	Passwd string "user passsword"
}
func main() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义

	u := reflect.TypeOf(user)
	fmt.Println(u.Align())
	fmt.Printf("%v\n",u)

	v := reflect.ValueOf(user)
	fmt.Printf("%v\n",v)

	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}