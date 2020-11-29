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
	ServiceName     string `json:"SERVICE_NAME" default:"omb-service"`
	RPCPort         int    `json:"RPC_PORT" default:"3000"`
	RestPort        int    `json:"REST_PORT" default:"4000"`
	OMDBHost        string `json:"OMDB_HOST" default:"https://www.omdbapi.com/"`
	OMDBAPIKey      string `json:"OMDB_API_KEY" default:"faf7e5bb"`
	EnableDBLogging bool   `json:"ENABLE_DB_LOGGING" default="true"`
	LogDBHost       string `json:"LOG_DB_HOST" default:"localhost:3306/test"`
	LogDBUser       string `json:"LOG_DB_USER" default="test"`
	LogDBPassword   string `json:"LOG_DB_PASSWORD" default="test"`
}

func New() *Config {
	once.Do(func() {
		instance = new(Config)
		envconfig.MustProcess("", instance)
	})
	return instance
}
