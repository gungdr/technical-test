package logger

import "log"

type dbLogger struct {
	serviceName string
}

func NewDBLogger(serviceName string) Logger {
	return &dbLogger{
		serviceName: serviceName,
	}
}

func (d dbLogger) Log(args ...interface{}) {
	args = append([]interface{}{d.serviceName}, args...)
	log.Println(args...)
}
