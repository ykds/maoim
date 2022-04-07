package comet

import (
	"encoding/json"
	"errors"
	"maoim/api/comet"
	"maoim/api/protocal"
	"maoim/pkg/websocket"
	"sync"
)

type Channel struct {
	IP   string
	Port string
	Key  string
	Conn *websocket.Conn

	signal chan *comet.PushMsg

	ProtoRing *Ring

	mu sync.RWMutex
}

func NewChannel(conn *websocket.Conn) *Channel {
	return &Channel{
		Conn:      conn,
		signal:    make(chan *comet.PushMsg, 10),
		ProtoRing: New(5),
	}
}

func (c *Channel) Ready() *comet.PushMsg {
	return <-c.signal
}

func (c *Channel) Push(p *comet.PushMsg) (err error) {
	select {
	case c.signal <- p:
	default:
		err = errors.New("signal is full")
	}
	return
}

func (c *Channel) ReadMessage(p *protocal.Proto) (err error) {
	op, payload, err := c.Conn.ReadWebSocket()
	if op == websocket.PingFrame {
		p = &protocal.Proto{Op: protocal.OpHeartBeat}
		return
	}
	return json.Unmarshal(payload, p)
}

func (c *Channel) WriteMessage(p *comet.PushMsg) error {
	marshal, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return c.Conn.WriteWebsocket(websocket.TextFrame, marshal)
}

func (c *Channel) Close() error {
	c.signal <- comet.ProtoFinish
	return c.Conn.Close()
}
