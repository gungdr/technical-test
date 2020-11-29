package database

import "github.com/jmoiron/sqlx"

const (
	READ  = 0
	WRITE = 1
)

type Database interface {
	Start() error
	Stop() error
	GetActiveDB(mode int) *sqlx.DB
}
