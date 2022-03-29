package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"maoim/internal/comet"
	"maoim/pkg/redis"
	"math/rand"
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

	engine := gin.Default()
	engine.GET("/", s.WsHandler)
	engine.POST("/register", s.Register)
	engine.POST("/login", s.Login)

	_ = engine.Run(":8000")
}




