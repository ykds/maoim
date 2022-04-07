package main

import (
	"context"
	"flag"
	"maoim/internal/user"
	"maoim/internal/user/grpc"
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

	dao := user.NewDao(r)
	service := user.NewService(dao)
	api := user.NewApi(service)
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
