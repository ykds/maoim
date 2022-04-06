package websocket

import (
	"bufio"
	"net/http"
)

type Request struct {
	Method     string
	RequestURI string
	Proto      string
	Host       string
	Header     http.Header

	reader *bufio.Reader
}
