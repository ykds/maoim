package comet

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"maoim/api/protocal"
	"maoim/pkg/websocket"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	HeartBeatInterval = 1 * time.Minute
)

func (s *Server) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{"code": 400, "message": "username或password为空"})
		return
	}

	user, err := s.LoadUser(username)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if user.ID != 0 {
		c.JSON(200, gin.H{"code": 400, "message": "用户名已被占用"})
		return
	}

	u := &User{
		ID: rand.Int63(),
		Username: username,
		Password: password,
	}
	err = s.SaveUser(u)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "register success", "data": u.ID})
}

func (s *Server) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{"code": 400, "message": "userId或password为空"})
		return
	}

	user, err := s.LoadUser(username)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if password == user.Password {
		cookie := map[string]string{"username": user.Username}
		ck, err := json.Marshal(cookie)
		if err != nil {
			c.JSON(200, gin.H{"code": 500, "message": err.Error()})
			return
		}
		c.JSON(200, gin.H{"code": 200, "message": "login success", "data": base64.StdEncoding.EncodeToString(ck)})
		return
	}
	c.JSON(200, gin.H{"code": 401, "message": "密码错误"})
}

func (s *Server) WsHandler(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request)
	if err != nil {
		return
	}

	user, err := s.auth(c.Request)
	if err != nil {
		_ = conn.WriteWebsocket(websocket.TextFrame, []byte("缺少cookie"))
		_ = conn.WriteWebsocket(websocket.CloseFrame, []byte(""))
		_ = conn.Close()
		return
	}

	go s.serveWebsocket(conn, user)
}

func (s *Server) auth(r *http.Request) (*User, error) {
	cookie := r.Header.Get("Cookie")
	if cookie == "" {
		return nil, errors.New("no cookie")
	}
	decodeString, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return nil, err
	}
	data := map[string]string{}
	err = json.Unmarshal(decodeString, &data)
	if err != nil {
		return nil, err
	}
	return s.LoadUser(data["username"])
}

func (s *Server) serveWebsocket(conn *websocket.Conn, user *User) {
	var (
		err error
		c = conn.GetConn().(net.Conn)
		hb = make(chan struct{})
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
	_ = s.bucket.PutChannel(ch.Key, ch)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return s.distributeMsg(ctx, ch)
	})
	g.Go(func() error {
		return s.heartbeat(ctx, hb)
	})
	g.Go(func() error {
		return s.ReadMessage(ctx, ch, hb)
	})
	if err = g.Wait(); err != nil {
		ch.Close()
	}
}

func (s *Server) ReadMessage(ctx context.Context, ch *Channel, hb chan<- struct{}) error {
	var lastHB = time.Now()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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
}

func (s *Server) distributeMsg(ctx context.Context, ch *Channel) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			p := ch.Ready()
			err := ch.WriteMessage(p)
			if err != nil {
				return err
			}
		}
	}
}

func (s *Server) heartbeat(ctx context.Context, hb <-chan struct{}) error {
	t := time.NewTicker(HeartBeatInterval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			return errors.New("heartbeat time out. connection closed")
		case <-ctx.Done():
			return ctx.Err()
		case <-hb:
			t.Reset(HeartBeatInterval)
		}
	}
}
