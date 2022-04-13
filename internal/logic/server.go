package logic

import (
	"github.com/gin-gonic/gin"
	"maoim/internal/logic/message"
	user2 "maoim/internal/logic/user"
	"net/http"
)

type Server struct {
	g     *gin.Engine
	srv   *http.Server
	messApi *message.Api
	userApi *user2.Api
}

func New(messApi *message.Api, userApi *user2.Api, g *gin.Engine) *Server {
	server := &http.Server{
		Addr:    ":8001",
		Handler: g,
	}

	s := &Server{
		g:     g,
		srv:   server,
		messApi: messApi,
		userApi: userApi,
	}
	return s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}
