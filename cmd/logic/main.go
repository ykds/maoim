package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"maoim/internal/logic/wire"
	comet2 "maoim/internal/pkg/grpc/comet"
	"maoim/pkg/redis"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "config file path", "config.json", "config file path")
	flag.Parse()

	c, err := redis.Load(filepath)
	if err != nil {
		panic(err)
	}
	r := redis.New(c)

	cometClient, err := comet2.NewCometGrpcClient()
	if err != nil {
		panic(err)
	}

	server := wire.Init(r, gin.Default(), cometClient)
	if err := server.Start(); err != nil {
		log.Println(err)
	}
}

