package database

import (
	"fmt"
	"log"
	"omdb/config"

	"github.com/jmoiron/sqlx"
)

type mysqlDatabase struct {
	master *sqlx.DB
}

func NewMysqlDabatabase(config *config.Config) Database {
	if config.EnableDBLogging == false {
		return &mysqlDatabase{}
	}
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@%s", config.LogDBUser, config.LogDBPassword, config.LogDBHost))
	if err != nil {
		log.Fatalln(err)
	}
	return &mysqlDatabase{
		master: db,
	}
}

func (db *mysqlDatabase) Start() error {
	return nil
}

func (db *mysqlDatabase) Stop() error {
	return nil
}

func (db *mysqlDatabase) GetActiveDB(mode int) *sqlx.DB {
	switch mode {
	case WRITE:
		return db.master
	default:
		return db.master
	}
}
