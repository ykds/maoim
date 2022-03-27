package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

func main() {
	ws, err := websocket.Dial("ws://127.0.0.1:8000/", "", "*")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("handshake complete")
	}

	err = websocket.Message.Send(ws, []byte("hello"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client send success")
	}

	var b string
	err = websocket.Message.Receive(ws, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client receive, " + b)
	}


}