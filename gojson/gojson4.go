package main

import (
	"fmt"
	"encoding/json"
)

type MenuState struct {
	Opened   string `json:"opened"`
	Selected string `json:"selected"`
}

type MenuNode struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	State    MenuState `json:"state"`
	Children []MenuNode `json:"children"`
}

func main() {
	var text = `
		{"id":"0","text":"菜单管理","state":{"opened":"true","selected":"true"},"children":[{"id":"2","text":"数据中心","state":{"opened":"true","selected":"true"},"children":null}]}
	`

	bytes := []byte(text)

	// Unmarshal string into structs.
	var menuNode MenuNode
	json.Unmarshal(bytes, &menuNode)

	// Loop over structs and display them.
	fmt.Printf("Id = %v, Text = %v , State= %v", menuNode.Id, menuNode.Text, menuNode.State.Opened)

	myNode := MenuNode{
		Id:  "0",
		Text: "菜单管理",
		State: MenuState{
			Opened:"true",
			Selected:"true",
		},
		Children:[]MenuNode{
			MenuNode{
				Id:  "2",
				Text: "数据中心",
				State: MenuState{
					Opened:"true",
					Selected:"true",
				},
			},
		},
	}
	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(myNode)
	// Convert bytes to string.
	s := string(b)
	fmt.Println(s)

}