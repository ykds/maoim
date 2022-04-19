package main

import (
	"context"
	"flag"
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
	flag.StringVar(&filepath, "config file path", "config.yaml", "config file path")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	config := conf.Load(filepath)
	r := redis.New(config.Redis)
	m := mysql.New(config.Mysql)

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

	time.Sleep(time.Second)
	go initlogic(config, m, r, engine)
	go initcomet(r, done, engine)

	<-sig
	httpServer.Shutdown(context.Background())
	close(done)
}

func initcomet(r *redis.Redis, done <-chan struct{}, g *gin.Engine) {
	s := comet.NewServer(r)
	grpcServer := grpc.New(s)

	g.GET("/", s.WsHandler)

	<-done

	grpcServer.GracefulStop()
}

func initlogic(c *conf.Config, m *mysql.Mysql, r *redis.Redis, g *gin.Engine) {
	server := wire.Init(c, r, m, g)
	if err := server.Start(); err != nil {
		server.Stop()
		log.Println(err)
	}
}
