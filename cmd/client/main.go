package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
	"strings"
)

func main() {
	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	//go client1(wg)
	//go client2(wg)
	//wg.Wait()

	config, err := websocket.NewConfig("ws://whg85s.natappfree.cc/", "*")
	if err != nil {
		panic(err)
	}
	config.Header["Cookie"] = []string{"{\"userId\":4639730689396572890}"}
	ws, err := websocket.DialConfig(config)
	if err != nil {
		panic(err)
	}

	read(ws)
	go write(ws)
}

func read(ws *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		strings.Replace(text, "\n", "", -1)

		payload := map[string]interface{}{
			"Tos": []string{"772737620168600365"},
			"Msg": text,
			"Seq": 1,
		}
		data, _ := json.Marshal(payload)
		err := websocket.Message.Send(ws, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func write(ws *websocket.Conn) {
	var b string
	for {
		err := websocket.Message.Receive(ws, &b)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			continue
		}
		msg := map[string]interface{}{}
		err = json.Unmarshal([]byte(b), &msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s: %s\n", msg["From"], msg["Msg"])
		}
	}
}
