package main

import (
	"fmt"
	"maoim/pkg/websocket"
	"net/http"
)

func main() {
	//http.Handle("/", websocket.Handler(connect))
	http.HandleFunc("/", connect)
	_ = http.ListenAndServe(":8000", nil)
}

func connect(w http.ResponseWriter, r *http.Request) {
	//var p string
	//for {
	//	err := websocket.Message.Receive(ws, &p)
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(p)
	//
	//	websocket.Message.Send(ws, "fuck")
	//}

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		return
	}

	for {
		read, op, data, err := conn.Read()
		if err != nil {
			conn.Write(websocket.TextFrame, []byte("fuck"))
		} else {
			//conn.Write(websocket.TextFrame, []byte("fuck"))
			s := fmt.Sprintf("read: %v, op: %d, data: %s", read, op, string(data))
			conn.Write(websocket.TextFrame, []byte(s))
		}
	}
}


