package main

import (
	"context"
	"flag"
	user2 "maoim/internal/logic/user"
	"maoim/internal/logic/user/grpc"
	"maoim/pkg/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "filepath", "config.json", "config file")
	flag.Parse()

	c, err := redis.Load(filepath)
	if err != nil {
		panic(err)
	}
	r := redis.New(c)

	dao := user2.NewDao(r)
	service := user2.NewService(dao)
	api := user2.NewApi(service)
	go api.Start()
	server := grpc.NewUserGrpcServer(service)

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGKILL)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	api.Shutdown(ctx)
	server.GracefulStop()
}
