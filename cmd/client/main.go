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
	config.Header["Cookie"] = []string{"{\"userId\":\"A\"}"}
	ws, err := websocket.DialConfig(config)
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)
	payload := map[string]interface{}{
		"From": "A",
		"To": "B",
		"Msg": "hello B",
	}
	data, _ := json.Marshal(payload)
	err = websocket.Message.Send(ws, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client A send success")
	}

	var b string
	err = websocket.Message.Receive(ws, &b)
	if err != nil {
		if err == io.EOF {
			wg.Done()
		}
		fmt.Println(err)
	} else {
		fmt.Println("client A receive, " + b)
	}

}

func client2(wg *sync.WaitGroup) {
	config, err := websocket.NewConfig("ws://127.0.0.1:8000/", "*")
	if err != nil {
		panic(err)
	}
	config.Header["Cookie"] = []string{"{\"userId\":\"B\"}"}
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
		fmt.Println("client B receive, " + b)

		payload := map[string]interface{}{
			"From": "B",
			"To": "A",
			"Msg": "hello A",
		}
		data, _ := json.Marshal(payload)
		err = websocket.Message.Send(ws, data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("client B send success")
		}
	}
}