package main

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Port string `yaml:"port"`
}

func LoadConfig() *Config {

	nv := "CONFIG_PATH_SIPLEWEBAPP"
	path := os.Getenv(nv)
	if path == "" {
		log.Fatalf("%s is not set", nv)
	}

	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s : %s", path, err.Error())
	}

	return &cfg
}
