package logic

import (
	"context"
	"github.com/gin-gonic/gin"
	"maoim/internal/logic/conf"
	"maoim/internal/logic/message"
	user2 "maoim/internal/logic/user"
	"net/http"
)

type Server struct {
	c *conf.Config
	g       *gin.Engine
	srv     *http.Server
	messApi *message.Api
	userApi *user2.Api
}

func New(c *conf.Config, messApi *message.Api, userApi *user2.Api, g *gin.Engine) *Server {
	server := &http.Server{
		Addr:    ":" + c.Logic.Port,
		Handler: g,
	}

	s := &Server{
		c: c,
		g:       g,
		srv:     server,
		messApi: messApi,
		userApi: userApi,
	}
	return s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	s.srv.Shutdown(context.Background())
	return s.userApi.Shutdown()
}
