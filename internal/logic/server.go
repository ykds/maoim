package logic

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "maoim/api/comet"
	"maoim/api/protocal"
	"maoim/internal/user"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	g *gin.Engine
	srv *http.Server
	comet pb.CometClient
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
	s := &Server{
		g: engine,
		srv: server,
		comet: cometClient,
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

func (s *Server) initRouter() {
	group := s.g.Group("/api")
	group.POST("/pushMsg", s.PushMsg)
}

func (s *Server) PushMsg(c *gin.Context) {
	var (
		arg struct {
			Keys []string `json:"keys"`
			Op int32 `json:"op"`
			Seq int32 `json:"seq"`
			Body string `json:"body"`
		}
	)
	err := c.BindJSON(&arg)
	if err != nil {
		log.Println(arg)
		c.JSON(400, gin.H{"code": 400, "message": "参数格式错误"})
		return
	}

	u, exists := c.Get("user")
	if !exists {
		log.Println(arg)
		c.JSON(401, gin.H{"code": 401, "message": "no auth"})
		return
	}
	us, _ := u.(*user.User)
	req := &pb.PushMsgReq{
		Keys: arg.Keys,
		PushMsg: &pb.PushMsg{
			FromKey: strconv.FormatInt(us.ID, 10),
			FromWho: us.Username,
			Proto: &protocal.Proto{
				Op: arg.Op,
				Seq: arg.Seq,
				Body: arg.Body,
			},
		},
	}

	_, err = s.comet.PushMsg(context.Background(), req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"code": 500, "message": "Internal Error"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "message": "success"})
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}