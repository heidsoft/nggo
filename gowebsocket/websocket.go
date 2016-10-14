package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
	"fmt"
)

func echoHandler(ws *websocket.Conn) {
	fmt.Println("websocket,消息处理")
	msg := make([]byte,0)
	ws.Read(msg)
	fmt.Println(string(msg))
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}