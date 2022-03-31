package comet

import (
	"encoding/json"
	"errors"
	"io"
	"maoim/api/protocal"
	"maoim/pkg/websocket"
	"sync"
)

type Channel struct {
	IP string
	Port string
	Key string
	Conn *websocket.Conn

	signals chan *protocal.Proto

	ProtoRing *Ring

	mu sync.RWMutex
}

func NewChannel(conn *websocket.Conn) *Channel {
	return &Channel{Conn: conn}
}

func (c *Channel) ReadMessage() (op int, p *Protocal, err error) {
	conn := c.Conn

	_, op, payload, err := conn.ReadWebSocket()
	if err != nil {
		return op, nil, err
	}

	switch op {
	case websocket.TextFrame, websocket.BinaryFrame:
		p, err := ParseMessage(payload)
		return op, p, err
	case websocket.PingFrame:
		err = conn.WriteWebsocket(websocket.PongFrame, payload)
		return
	case websocket.CloseFrame:
		_ = conn.Close()
		return op, nil, io.EOF
	}
	return op, nil, errors.New("不支持的操作")
}

func (c *Channel) WriteMessage(p *Protocal) error {
	data := make(map[string]interface{}, 2)
	data["FromId"] = p.FromId
	data["From"] = p.From
	data["Msg"] = p.Msg
	data["Seq"] = p.Seq

	marshal, _ := json.Marshal(data)
	return c.Conn.WriteWebsocket(websocket.TextFrame, marshal)
}

func ParseMessage(payload []byte) (p *Protocal, err error) {
	p = &Protocal{}
	err = json.Unmarshal(payload, p)
	return
}