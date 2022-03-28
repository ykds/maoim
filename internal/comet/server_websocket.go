package comet

import (
	"encoding/json"
	"errors"
	"maoim/pkg/websocket"
	"net"
	"net/http"
)


type Message struct {
	From string
	To string
	Msg string
	MsgType int
}

func (s *Server) WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		return
	}

	userId, err := s.auth(r)
	if err != nil {
		conn.WriteWebsocket(websocket.TextFrame, []byte("缺少cookie"))
		conn.WriteWebsocket(websocket.CloseFrame, []byte(""))
		return
	}

	go s.serveWebsocket(conn, userId)
}


func (s *Server) auth(r *http.Request) (string, error) {
	cookie := r.Header.Get("Cookie")
	if cookie == "" {
		return "", errors.New("no cookie")
	}
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(cookie), &data)
	if err != nil {
		return "", err
	}
	return data["userId"].(string), nil
}


func (s *Server) serveWebsocket(conn *websocket.Conn, userId string) {
	var err error
	c := conn.GetConn().(net.Conn)

	ch := NewChannel(conn)
	ch.IP, ch.Port, err = net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		ch.Conn.WriteWebsocket(websocket.TextFrame, []byte("IP Address format error"))
		conn.WriteWebsocket(websocket.CloseFrame, []byte(""))
		return
	}
	ch.Key = userId
	_ = s.bucket.PutChannel(ch.Key, ch)

	for {
		ch.ReadMessage()
	}
}

