package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"maoim/internal/comet"
	"maoim/internal/comet/grpc"
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

	time.Sleep(time.Second)
	go initlogic(r, engine)
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

func initlogic(r *redis.Redis, g *gin.Engine) {
	mysqlConfig := mysql.Default()
	mysqlConfig.DbName = "maoim"
	m := mysql.New(mysqlConfig)

	server := wire.Init(r, m, g)
	if err := server.Start(); err != nil {
		server.Stop()
		log.Println(err)
	}
}
