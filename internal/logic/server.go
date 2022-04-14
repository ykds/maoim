package logic

import (
	"context"
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
	messApi.InitRouter(g)
	userApi.InitRouter(g)
	return s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	s.srv.Shutdown(context.Background())
	return s.userApi.Shutdown()
}
