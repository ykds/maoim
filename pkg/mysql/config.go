package mysql

import "maoim/pkg/yaml"

var Conf *Config

type Config struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DbName string `json:"db_name" yaml:"db_name"`
}

func Init(confPath string) error {
	Conf = Default()
	return yaml.DecodeFile(confPath, Conf)
}

func Default() *Config {
	return &Config{
		Host: "127.0.0.1",
		Port: "3306",
		Username: "root",
		Password: "123456",
		DbName: "",
	}
}