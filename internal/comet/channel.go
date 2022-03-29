package comet

import (
	"encoding/json"
	"errors"
	"io"
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
		return nil, io.EOF
	}
	return nil, errors.New("不支持的操作")
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