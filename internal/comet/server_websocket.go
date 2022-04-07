package comet

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"maoim/api/comet"
	"maoim/api/protocal"
	pb "maoim/api/user"
	"maoim/internal/pkg/utils"
	"maoim/internal/user"
	"maoim/pkg/websocket"
	"net"
	"strconv"
	"time"
)

const (
	HeartBeatInterval = 5 * time.Minute
)

func (s *Server) auth(c *gin.Context) (*user.User, error) {
	token := c.Request.Header.Get("token")
	if token == "" {
		return nil, fmt.Errorf("缺少token")
	}

	username, err := utils.ValidToken(token)
	if err != nil {
		return nil, fmt.Errorf("token错误")
	}
	userReply, err := s.userClient.GetUserByUsername(context.Background(), &pb.GetUserReq{Username: username})
	if err != nil {
		return nil, fmt.Errorf("token错误")
	}
	userId, err := strconv.ParseInt(userReply.Id, 10, 64)
	u := &user.User{
		ID: userId,
		Username: userReply.Username,
		Password: userReply.Password,
	}
	return u, nil
}

func (s *Server) WsHandler(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request)
	if err != nil {
		return
	}

	u, err := s.auth(c)
	if err != nil {
		_ = conn.WriteWebsocket(websocket.TextFrame, []byte(err.Error()))
		_ = conn.WriteWebsocket(websocket.CloseFrame, []byte(""))
		_ = conn.Close()
		return
	}

	go s.serveWebsocket(conn, u)
}

func (s *Server) serveWebsocket(conn *websocket.Conn, user *user.User) {
	var (
		err error
		c   = conn.GetConn().(net.Conn)
		hb  = make(chan struct{})
	)

	fmt.Printf("%d is online.\n", user.ID)

	ch := NewChannel(conn)
	ch.IP, ch.Port, err = net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		_ = ch.Conn.WriteWebsocket(websocket.TextFrame, []byte("IP Address format error"))
		_ = conn.WriteWebsocket(websocket.CloseFrame, []byte(""))
		return
	}
	ch.Key = strconv.FormatInt(user.ID, 10)
	_ = s.bucket.PutChannel(user.Username, ch)
	defer s.bucket.DeleteChannel(user.Username)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return s.ReadMessage(ctx, ch, hb)
	})
	g.Go(func() error {
		return s.distributeMsg(ctx, ch)
	})
	g.Go(func() error {
		if err = s.heartbeat(ctx, hb); err != nil {
			ch.Close()
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		fmt.Println(ch.Key + "is offline")
	}
}

func (s *Server) ReadMessage(ctx context.Context, ch *Channel, hb chan<- struct{}) error {
	var lastHB = time.Now()
	for {
		p := &protocal.Proto{}
		err := ch.ReadMessage(p)
		if err != nil {
			return err
		}
		if p.Op == protocal.OpHeartBeat {
			if now := time.Now(); now.Sub(lastHB) > HeartBeatInterval {
				// TODO refresh conn timeout
				hb <- struct{}{}
				lastHB = now
			}
		}
	}
}

func (s *Server) distributeMsg(ctx context.Context, ch *Channel) error {
	for {
		p := ch.Ready()
		if p == comet.ProtoFinish {
			return fmt.Errorf("finish")
		}
		err := ch.WriteMessage(p)
		if err != nil {
			return err
		}
	}
}

func (s *Server) heartbeat(ctx context.Context, hb <-chan struct{}) error {
	t := time.NewTicker(HeartBeatInterval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			return fmt.Errorf("heartbeat time out. connection closed")
		case <-hb:
			t.Reset(HeartBeatInterval)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
