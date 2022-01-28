package config

import (
	"github.com/divilla/golastore/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		Dsn string `yaml:"dsn"`
	}
)

func New(log *logger.Logger) *Config {
	var cfg Config
	out, err := os.ReadFile("config/local.yaml")
	if err != nil {
		log.ErrorWithStack(err)
	}

	if err = yaml.Unmarshal(out, &cfg); err != nil {
		log.ErrorWithStack(err)
	}

	return &cfg
}
