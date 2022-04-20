package comet

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"maoim/api/comet"
	mpb "maoim/api/message"
	"maoim/api/protocal"
	pb "maoim/api/user"
	user2 "maoim/internal/logic/user"
	"maoim/internal/pkg/utils"
	"maoim/pkg/websocket"
	"net"
	"time"
)

const (
	HeartBeatInterval = 100 * time.Minute
)

func (s *Server) auth(c *gin.Context) (*user2.User, error) {
	token := c.Request.Header.Get("token")
	if token == "" {
		return nil, fmt.Errorf("缺少token")
	}

	userId, username, err := utils.ValidToken(token)
	if err != nil {
		return nil, fmt.Errorf("token错误")
	}
	reply, err := s.userClient.Connect(context.Background(), &pb.ConnectReq{
		UserId:   userId,
		Username: username,
	})
	if err != nil {
		return nil, fmt.Errorf("连接异常")
	}
	u := &user2.User{}
	u.ID = reply.UserId
	u.Username = reply.UserName
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

func (s *Server) serveWebsocket(conn *websocket.Conn, user *user2.User) {
	var (
		err error
		c   = conn.GetConn().(net.Conn)
		hb  = make(chan struct{})
	)

	fmt.Printf("(%s, %s) is online.\n", user.ID, user.Username)

	ch := NewChannel(conn)
	ch.IP, ch.Port, err = net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		_ = ch.Conn.WriteWebsocket(websocket.TextFrame, []byte("IP Address format merror"))
		_ = conn.WriteWebsocket(websocket.CloseFrame, []byte(""))
		return
	}
	ch.Key = user.ID
	_ = s.bucket.PutChannel(ch.Key, ch)
	defer s.bucket.DeleteChannel(ch.Key)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return s.ReadMessage(ch, hb)
	})
	g.Go(func() error {
		return s.distributeMsg(ch)
	})
	g.Go(func() error {
		if err = s.heartbeat(ctx, hb, ch.Key); err != nil {
			_ = ch.Close()
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		fmt.Printf("(%s, %s) is offline\n", user.ID, user.Username)
	}
}

type AckReq struct {
	MsgIdList []string `json:"msg_id_list"`
}

func (s *Server) ReadMessage(ch *Channel, hb chan<- struct{}) error {
	var lastHB = time.Now()
	for {
		p := &protocal.Proto{}
		err := ch.ReadMessage(p)
		if err != nil {
			return err
		}
		if p.Op == protocal.OpHeartBeat {
			if now := time.Now(); now.Sub(lastHB) > HeartBeatInterval {
				_ = ch.Push(comet.ProtoHeartBeatReply)
				hb <- struct{}{}
				lastHB = now
			}
		} else if p.Op == protocal.OpAck {
			ackReq := &AckReq{}
			err = json.Unmarshal(p.Body, ackReq)
			if err != nil {
				fmt.Println(err)
				_ = ch.Push(&protocal.Proto{
					Op:   protocal.OpErr,
					Body: []byte("ack包解析错误"),
				})
				continue
			}
			if len(ackReq.MsgIdList) == 0 {
				continue
			}
			_, err := s.messageClient.AckMsg(context.Background(), &mpb.AckReq{
				UserId:   ch.Key,
				MsgId:    ackReq.MsgIdList,
			})
			if err != nil {
				fmt.Println(err)
				_ = ch.Push(&protocal.Proto{
					Op:   protocal.OpErr,
					Body: []byte("ack消息失败"),
				})
				continue
			}
		}
	}
}

func (s *Server) distributeMsg(ch *Channel) error {
	for {
		p := ch.Ready()
		switch p {
		case comet.ProtoFinish:
			return fmt.Errorf("finish")
		case comet.ProtoHeartBeatReply:
			if err := ch.WriteHeartBeat(); err != nil {
				return err
			}
			continue
		default:
			err := ch.WriteMessage(p)
			if err != nil {
				return err
			}
		}
	}
}

func (s *Server) heartbeat(ctx context.Context, hb <-chan struct{}, key string) error {
	t := time.NewTicker(HeartBeatInterval)
	defer func() {
		t.Stop()
		_, _ = s.userClient.Disconnect(context.Background(), &pb.DisconnectReq{
			UserId:   key,
		})
	}()

	for {
		select {
		case <-hb:
			t.Reset(HeartBeatInterval)
			_, _ = s.userClient.Connect(context.Background(), &pb.ConnectReq{
				UserId:   key,
			})
		case <-t.C:
			return fmt.Errorf("heartbeat time out. connection closed")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
