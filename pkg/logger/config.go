package logger

import "maoim/pkg/yaml"

type Config struct {
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"max_size" yaml:"max_size"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups"`
	MaxAge     int    `json:"max_age" yaml:"max_age"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

func Default() *Config {
	return &Config{
		Filename:   "logs/err.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   false,
	}
}

func Init(configFile string) *Config {
	conf := Default()
	err := yaml.DecodeFile(configFile, conf)
	if err != nil {
		panic(err)
	}
	return conf
}
