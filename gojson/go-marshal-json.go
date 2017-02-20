package main

import (
	"encoding/json"
	"time"
	"fmt"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	fmt.Println("转换json 格式......")
	return json.Marshal(&struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastSeen string  `json:"lastSeen"`
	}{
		ID:       u.ID,
		Name:     u.Name,
		LastSeen: "123456",
	})
}

func main() {
	user := new(MyUser)
	user.ID=1
	user.Name="test"
	user.LastSeen = time.Now()
	fmt.Println(user)

	userByte,err := json.Marshal(user)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(userByte))
}