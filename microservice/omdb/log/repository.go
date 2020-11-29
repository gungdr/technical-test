package log

import (
	"context"
	"log"
	"omdb/config"
	"omdb/database"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	config *config.Config
	db     database.Database
}

type Repository interface {
	Insert(ctx context.Context, data Log) error
}

const (
	insertStatement = `
		INSERT INTO log (service_name,message,time)
		VALUES(:service_name,:message,:time)
	`
)

func NewRepository(conf *config.Config, db database.Database) Repository {
	return &repository{
		config: conf,
		db:     db,
	}
}

func (repo *repository) Insert(ctx context.Context, data Log) error {
	_, err := repo.db.GetActiveDB(database.WRITE).MustBegin().NamedExecContext(ctx,
		insertStatement,
		data,
	)
	if err != nil {
		log.Println("Repos Insert:", err)
		return err
	}
	return nil
}
