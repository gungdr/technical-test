package logger

import (
	"context"
	"log"
	"omdb/config"
	logdb "omdb/log"
	"time"

	"encoding/json"
)

type dbLogger struct {
	config      *config.Config
	serviceName string
	logRepo     logdb.Repository
}

func NewDBLogger(config *config.Config, serviceName string, logRepo logdb.Repository) Logger {
	return &dbLogger{
		config:      config,
		serviceName: serviceName,
		logRepo:     logRepo,
	}
}

func (d dbLogger) Log(args ...interface{}) {
	args = append([]interface{}{d.serviceName}, args...)
	log.Println(args...)
	message, err := json.Marshal(args)
	if err != nil {
		log.Println("Log Error:", err)
		return
	}
	if d.config.EnableDBLogging {
		go d.logRepo.Insert(context.Background(), logdb.Log{
			ServiceName: d.serviceName,
			Message:     string(message),
			Time:        time.Now(),
		})
	}
}
