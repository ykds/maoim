package comet

import (
	"encoding/json"
	"errors"
	"maoim/pkg/websocket"
	"sync"
)

type Channel struct {
	IP string
	Port string
	Key string

	Conn *websocket.Conn

	mu sync.RWMutex

	Seq int
}

func NewChannel(conn *websocket.Conn) *Channel {
	return &Channel{Conn: conn, Seq: 0}
}

func (c *Channel) incrSeq() {
	c.Seq++
}


func (c *Channel) ReadMessage() (p *Protocal, err error) {
	conn := c.Conn

	_, op, payload, err := conn.ReadWebSocket()
	if err != nil {
		return nil, err
	}

	switch op {
	case websocket.TextFrame, websocket.BinaryFrame:
		return ParseMessage(payload)
	case websocket.PingFrame:
		// TODO handle pong
	case websocket.CloseFrame:
		_ = conn.Close()
		return
	}
	return nil, errors.New("不支持的操作")
}

func (c *Channel) WriteMessage(p *Protocal) error {
	return nil
}

func ParseMessage(payload []byte) (p *Protocal, err error) {
	err = json.Unmarshal(payload, p)
	return
}