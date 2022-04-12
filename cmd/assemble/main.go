package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"maoim/internal/comet"
	"maoim/internal/comet/grpc"
	"maoim/internal/logic"
	user2 "maoim/internal/logic/user"
	ugrpc "maoim/internal/logic/user/grpc"
	"maoim/pkg/redis"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "config file path", "config.json", "config file path")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	c, err := redis.Load(filepath)
	if err != nil {
		panic(err)
	}
	r := redis.New(c)

	done := make(chan struct{})
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	engine := gin.Default()

	httpServer := http.Server{
		Addr:    ":8000",
		Handler: engine,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	go inituser(r, done, engine)
	time.Sleep(time.Second)
	go initlogic(engine)
	go initcomet(r, done, engine)

	<-sig
	httpServer.Shutdown(context.Background())
	close(done)
}

func initcomet(r *redis.Redis, done <-chan struct{}, g *gin.Engine)  {
	s := comet.NewServer(r)
	grpcServer := grpc.New(s)

	g.GET("/", s.WsHandler)

	<-done

	grpcServer.GracefulStop()
}

func initlogic(g *gin.Engine) {
	_ = logic.NewDebug(g)
}

func inituser(r *redis.Redis, done <-chan struct{}, g *gin.Engine)  {
	dao := user2.NewDao(r)
	service := user2.NewService(dao)
	_ = user2.NewApiDebug(service, g)
	server := ugrpc.NewUserGrpcServer(service)
	<-done
	server.GracefulStop()
}