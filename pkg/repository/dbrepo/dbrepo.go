package dbrepo

import (
	"database/sql"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/repository"
)

var app *config.AppConfig

type mysqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewMySQLRepo creates the repository
func NewMySQLRepo(Conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	app = a
	return &mysqlDBRepo{
		App: a,
		DB:  Conn,
	}
}
