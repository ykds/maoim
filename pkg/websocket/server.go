package websocket

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

const (
	_wsKeyMagicNum = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (*Conn, error) {
	if r.Method != "GET" {
		return nil, errors.New("http method is error")
	}

	if strings.ToLower(r.Header.Get("Upgrade")) != "websocket" {
		return nil, errors.New("upgrade is loss")
	}

	if !strings.Contains(r.Header.Get("Connection"), "Upgrade") {
		return nil, errors.New("connection lost upgrade")
	}

	if r.Header.Get("Sec-Websocket-Version") != "13" {
		return nil, errors.New("web")
	}

	wsKey := r.Header.Get("Sec-Websocket-Key")
	if wsKey == "" {
		return nil, errors.New("Sec-Websocket-Key is loss")
	}

	h, ok := w.(http.Hijacker)
	if !ok {
		return nil, errors.New("w is not implement Hiijacker")
	}

	conn, rwc, err := h.Hijack()
	if err != nil {
		return nil, err
	}

	_, _ = rwc.WriteString("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\n")
	_, _ = rwc.WriteString("Sec-Websocket-Accept: " + computeAcceptKey(wsKey) + "\r\n")
	_, _ = rwc.WriteString("Sec-Websocket-Version: 13\r\n\r\n")

	if err  = rwc.Flush(); err != nil {
		return nil, err
	}

	return newConn(conn, rwc.Reader, rwc.Writer), nil
}

func computeAcceptKey(key string) string {
	hash := sha1.New()
	hash.Write([]byte(key + _wsKeyMagicNum))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func maskBytes(key []byte, data []byte) {
	pos := 0
	data[pos] ^= key[pos&3]
	pos++
}