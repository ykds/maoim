package comet

import (
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

	signal chan *protocal.Proto

	//ProtoRing *Ring

	mu sync.RWMutex
}

func NewChannel(conn *websocket.Conn) *Channel {
	return &Channel{
		Conn:      conn,
		signal:    make(chan *protocal.Proto, 10),
		//ProtoRing: New(5),
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
	op, payload, err := c.Conn.ReadWebSocket()
	if err != nil {
		return
	}
	if op == websocket.PingFrame {
		p = &protocal.Proto{Op: protocal.OpHeartBeat}
		return nil
	}
	p.Unpack(payload)
	return
}

func (c *Channel) WriteMessage(p *protocal.Proto) error {
	return c.Conn.WriteWebsocket(websocket.TextFrame, p.Pack())
}

func (c *Channel) WriteHeartBeat() error {
	return c.Conn.WriteWebsocket(websocket.PongFrame, (&protocal.Proto{}).PackHeartBeat())
}

func (c *Channel) Close() error {
	c.signal <- comet.ProtoFinish
	return c.Conn.Close()
}
