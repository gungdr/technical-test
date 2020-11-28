package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	Port       int    `json:"PORT" default:"3000"`
	OMDBHost   string `json:"OMDB_HOST" default:"https://www.omdbapi.com/"`
	OMDBAPIKey string `json:"OMDB_API_KEY" default:"faf7e5bb"`
}

func New() *Config {
	once.Do(func() {
		instance = new(Config)
		envconfig.MustProcess("", instance)
	})
	return instance
}
