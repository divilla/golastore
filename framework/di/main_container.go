package di

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"go.uber.org/zap"
)

type (
	MainContainer struct {
		logger *zap.Logger
		config *config.Config
	}
)

func InitMainContainer() *MainContainer {
	log := getLogger()
	conf := getConfig()

	return &MainContainer{
		logger: log,
		config: conf,
	}
}

func (m *MainContainer) Config() *config.Config {
	return m.config
}

func (m *MainContainer) Logger() *zap.Logger {
	return m.logger
}

func (m *MainContainer) Close() error {
	return nil
}

func getLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger
}

func getConfig() *config.Config {
	c := config.NewEmpty("spa")
	c.AddDriver(yaml.Driver)
	err := c.LoadFiles("config/local.yaml")
	if err != nil {
		panic(err)
	}

	return c
}
