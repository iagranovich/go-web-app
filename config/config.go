package config

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

func Load() *Config {

	ev := "CONFIG_PATH_SIMPLEWEBAPP"
	path := os.Getenv(ev)
	if path == "" {
		log.Fatalf("environment variable %s is not set", ev)
	}

	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s : %s", path, err.Error())
	}

	return &cfg
}
