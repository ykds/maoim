package main

import (
	"fmt"
	"maoim/pkg/websocket"
	"net/http"
)

func main() {
	http.HandleFunc("/", connect)
	_ = http.ListenAndServe(":8000", nil)
}

func connect(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		return
	}

	for {
		read, op, data, err := conn.Read()
		if err != nil {
			conn.Write(websocket.TextFrame, []byte("fuck"))
		} else {
			s := fmt.Sprintf("read: %v, op: %d, data: %s", read, op, string(data))
			conn.Write(websocket.TextFrame, []byte(s))
		}
	}
}


