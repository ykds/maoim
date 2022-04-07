package logic

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "maoim/api/comet"
	upb "maoim/api/user"
	"maoim/internal/user"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	g     *gin.Engine
	srv   *http.Server
	comet pb.CometClient
	user upb.UserClient
}

func New() *Server {
	engine := gin.Default()
	server := &http.Server{
		Addr:    ":8002",
		Handler: engine,
	}
	cometClient, err := newCometGrpcClient()
	if err != nil {
		panic(err)
	}
	userClient, err := newUserGrpcClient()
	if err != nil {
		panic(err)
	}

	s := &Server{
		g:     engine,
		srv:   server,
		comet: cometClient,
		user: userClient,
	}

	s.initRouter()
	return s
}

func newCometGrpcClient() (pb.CometClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dial, err := grpc.DialContext(ctx, ":8001", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		return nil, err
	}
	return pb.NewCometClient(dial), nil
}

func newUserGrpcClient() (upb.UserClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dial, err := grpc.DialContext(ctx, ":8003", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		return nil, err
	}
	return upb.NewUserClient(dial), nil
}

func (s *Server) auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Abort()
			return
		}
		reply, err := s.user.Auth(context.Background(), &upb.AuthReq{Token: token})
		if err != nil {
			c.Abort()
			return
		}
		userId, err := strconv.ParseInt(reply.Id, 10, 64)
		u := &user.User{
			ID: userId,
			Username: reply.Username,
			Password: reply.Password,
		}
		c.Set("user", u)
		c.Next()
	}
}

func (s *Server) initRouter() {
	group := s.g.Group("/api", s.auth())
	group.POST("/pushMsg", s.PushMsg)
}


func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}
