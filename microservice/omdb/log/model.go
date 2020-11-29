package log

import "time"

type Log struct {
	ServiceName string    `db:"service_name"`
	Message     string    `db:"message"`
	Time        time.Time `db:"time"`
}
