package comet

import (
	"encoding/json"
	"errors"
	"maoim/api/protocal"
	"maoim/pkg/websocket"
	"sync"
)

type Channel struct {
	IP string
	Port string
	Key string
	Conn *websocket.Conn

	signal chan *protocal.Proto

	ProtoRing *Ring

	mu sync.RWMutex
}

func NewChannel(conn *websocket.Conn) *Channel {
	return &Channel{
		Conn: conn,
		signal: make(chan *protocal.Proto, 10),
		ProtoRing: New(5),
	}
}

func (c *Channel) Ready() *protocal.Proto {
	return <-c.signal
}

func (c *Channel) Push(p *protocal.Proto) (err error) {
	select {
	case c.signal <- p:
	default:
		err = errors.New("signal is full")
	}
	return
}

func (c *Channel) ReadMessage(p *protocal.Proto) (err error) {
	payload, err := c.Conn.ReadWebSocket()
	return json.Unmarshal(payload, p)
}

func (c *Channel) WriteMessage(p *protocal.Proto) error {
	marshal, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return c.Conn.WriteWebsocket(websocket.TextFrame, marshal)
}

func (c *Channel) Close() error {
	return c.Conn.Close()
}