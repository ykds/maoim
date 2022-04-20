package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"maoim/internal/comet"
	"maoim/internal/comet/grpc"
	"maoim/internal/logic/conf"
	"maoim/internal/logic/wire"
	"maoim/pkg/mysql"
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
	flag.StringVar(&filepath, "config", "config.yaml", "config file path")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	config := conf.Load(filepath)
	r := redis.New(config.Redis)
	m := mysql.New(config.Mysql)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	engine := gin.Default()

	httpServer := http.Server{
		Addr:    ":" + config.Assemble.Port,
		Handler: engine,
	}
	go func() {
		fmt.Println("run on " + config.Assemble.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go initlogic(ctx, config, m, r, engine)
	go initcomet(ctx, r, engine)

	<-sig
	cancelFunc()
	httpServer.Shutdown(context.Background())
}

func initcomet(ctx context.Context, r *redis.Redis, g *gin.Engine) {
	s := comet.NewServer(r)
	grpcServer := grpc.New(s)

	g.GET("/", s.WsHandler)

	<-ctx.Done()

	grpcServer.GracefulStop()
}

func initlogic(ctx context.Context, c *conf.Config, m *mysql.Mysql, r *redis.Redis, g *gin.Engine) {
	server := wire.Init(c, r, m, g)
	go func() {
		if err := server.Start(); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	server.Stop()
}
