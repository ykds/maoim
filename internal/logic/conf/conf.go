package conf

import (
	"maoim/pkg/mysql"
	"maoim/pkg/redis"
	"maoim/pkg/yaml"
)

type Config struct {
	Mysql    *mysql.Config
	Redis    *redis.Config
	Logic    *Server
	Assemble *AssembleServer
}

type Server struct {
	Port string `json:"port" yaml:"port"`
}

type AssembleServer struct {
	Port string `json:"port" yaml:"port"`
}

func Load(filepath string) *Config {
	c := &Config{}
	err := yaml.DecodeFile(filepath, c)
	if err != nil {
		panic(err)
	}
	return c
}
