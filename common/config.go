package common

import (
	"errors"
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/consul"
)

func InitConfig() config.Config {
	// Initialize Consul as a source
	consulSource := consul.NewSource(consul.WithPrefix("/"))

	// Create new config
	conf := config.NewConfig()

	// Load source
	conf.Load(consulSource)

	return conf
}

func GetConfigIntByPath(conf config.Config, path ...string) int {
	setting := conf.Get(path...).Int(-1)

	if setting == -1 {
		panic(generateNewConfigError(path...))
	} else {
		return setting
	}
}

func GetConfigStringByPath(conf config.Config, path ...string) string {
	setting := conf.Get(path...).String("error")

	if setting == "error" {
		panic(generateNewConfigError(path...))
	} else {
		return setting
	}
}

func generateNewConfigError(path ...string) error {
	return errors.New(fmt.Sprintf("no config could be found for path: %v", path))
}
