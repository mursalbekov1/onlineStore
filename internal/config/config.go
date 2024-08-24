package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env  string     `yaml:"env"`
	Http HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Panic(err)
	}

	return &cfg
}
