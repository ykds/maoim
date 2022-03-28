package main

import (
	"maoim/internal/comet"
	"net/http"
)

func main() {
	s := comet.NewServer()

	http.HandleFunc("/", s.WsHandler)
	_ = http.ListenAndServe(":8000", nil)
}




