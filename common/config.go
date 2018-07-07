package common

import (
	"errors"
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/consul"
)

func GetConfigByPath(path ...string) string {
	conf := initConfig()
	return conf.Get(path...).String(generateNewConfigError(path...).Error())
}

func initConfig() config.Config {
	// Create Consul as a source
	consulSource := consul.NewSource(consul.WithPrefix("/"))

	// Create new config
	conf := config.NewConfig()

	// Load file source
	conf.Load(consulSource)

	return conf
}

func generateNewConfigError(path ...string) error {
	return errors.New(fmt.Sprintf("No config could be found for path '%v'", path))
}
