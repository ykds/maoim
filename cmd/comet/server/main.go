package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"maoim/internal/comet"
	"maoim/internal/comet/grpc"
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

	s := comet.NewServer(r)
	grpcServer := grpc.New(s)

	engine := gin.Default()
	engine.GET("/", s.WsHandler)
	engine.POST("/register", s.Register)
	engine.POST("/login", s.Login)

	httpServer := http.Server{
		Addr:    ":8000",
		Handler: engine,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGKILL)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
	grpcServer.GracefulStop()
}
