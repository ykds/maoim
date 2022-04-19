package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"maoim/internal/logic/conf"
	"maoim/internal/logic/wire"
	"maoim/pkg/mysql"
	"maoim/pkg/redis"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "config file path", "config.yaml", "config file path")
	flag.Parse()

	config := conf.Load(filepath)
	r := redis.New(config.Redis)
	m := mysql.New(config.Mysql)
	server := wire.Init(config, r, m, gin.Default())
	if err := server.Start(); err != nil {
		server.Stop()
		log.Println(err)
	}
}
