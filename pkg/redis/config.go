package redis

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func Load(file string) (*Config, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = json.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
