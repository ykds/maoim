package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go client1(wg)
	go client2(wg)
	wg.Wait()
}

func client1(wg *sync.WaitGroup) {
	config, err := websocket.NewConfig("ws://127.0.0.1:8000/", "*")
	if err != nil {
		panic(err)
	}
	config.Header["Cookie"] = []string{"{\"userId\":4639730689396572890}"}
	ws, err := websocket.DialConfig(config)
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)
	payload := map[string]interface{}{
		"Tos": []string{"8071612869869735209"},
		"Msg": "hello hxy",
		"Seq": 1,
	}
	data, _ := json.Marshal(payload)
	err = websocket.Message.Send(ws, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client dyk send success")
	}

	var b string
	err = websocket.Message.Receive(ws, &b)
	if err != nil {
		if err == io.EOF {
			wg.Done()
		}
		fmt.Println(err)
	} else {
		fmt.Println("client dyk receive, " + b)
	}

}

func client2(wg *sync.WaitGroup) {
	config, err := websocket.NewConfig("ws://127.0.0.1:8000/", "*")
	if err != nil {
		panic(err)
	}
	config.Header["Cookie"] = []string{"{\"userId\":8071612869869735209}"}
	ws, err := websocket.DialConfig(config)

	time.Sleep(2 * time.Second)

	var b string
	err = websocket.Message.Receive(ws, &b)
	if err != nil {
		if err == io.EOF {
			wg.Done()
		}
		fmt.Println(err)
	} else {
		fmt.Println("client hxy receive, " + b)

		payload := map[string]interface{}{
			"Tos": []string{"4639730689396572890"},
			"Msg": "hello dyk",
			"Seq": 1,
		}
		data, _ := json.Marshal(payload)
		err = websocket.Message.Send(ws, data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("client hxy send success")
		}
	}
}