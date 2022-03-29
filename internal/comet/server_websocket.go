package comet

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"maoim/pkg/websocket"
	"math/rand"
	"net"
	"net/http"
	"strconv"
)

func (s *Server) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{"code": 400, "message": "username或password为空"})
		return
	}
	u := &User{
		ID: rand.Int63(),
		Username: username,
		Password: password,
	}
	err := s.SaveUser(u)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "register success", "data": u.ID})
}

func (s *Server) Login(c *gin.Context) {
	userId := c.PostForm("userId")
	password := c.PostForm("password")
	if userId == "" || password == "" {
		c.JSON(200, gin.H{"code": 400, "message": "userId或password为空"})
		return
	}

	user, err := s.LoadUser(userId)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "message": err.Error()})
		return
	}
	if password == user.Password {
		cookie := map[string]int64{"userId": user.ID}
		ck, err := json.Marshal(cookie)
		if err != nil {
			c.JSON(200, gin.H{"code": 500, "message": err.Error()})
			return
		}
		c.JSON(200, gin.H{"code": 200, "message": "login success", "data": string(ck)})
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
	data := map[string]int64{}
	err := json.Unmarshal([]byte(cookie), &data)
	if err != nil {
		return nil, err
	}

	return s.LoadUser(strconv.FormatInt(data["userId"], 10))
}


func (s *Server) serveWebsocket(conn *websocket.Conn, user *User) {
	defer conn.Close()

	var err error
	c := conn.GetConn().(net.Conn)

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

	for {
		p, err := ch.ReadMessage()
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("%d is offline.\n", user.ID)
				return
			}
			log.Println(err)
			continue
		}
		p.From = user.Username
		_ = s.PushMsg(p)
	}
}

func (s *Server) PushMsg(p *Protocal) error {
	for _, key := range p.Tos {
		channel, err := s.Bucket().GetChannel(key)
		if err != nil {
			log.Println(err)
			continue
		}
		_ = channel.WriteMessage(p)
	}
	return nil
}
