package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"maoim/internal/logic/wire"
	"maoim/pkg/mysql"
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

	mysqlConfig := mysql.Default()
	mysqlConfig.DbName = "maoim"

	m := mysql.New(mysqlConfig)
	server := wire.Init(r, m, gin.Default())
	if err := server.Start(); err != nil {
		server.Stop()
		log.Println(err)
	}
}
