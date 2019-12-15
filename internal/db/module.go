package db

import (
	"database/sql"

	"github.com/3ammari/sample-fx-app/internal/env"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"go.uber.org/fx"
)

//Module Database Module
var Module = fx.Provide(New)

//Params  gets env config
type Params struct {
	fx.In
	Config env.Config
}

//New creates a db connection
func New(p Params) (*sql.DB, error) {
	db, err := sql.Open("mysql", p.Config.DatabaseURL)
	return db, err
}
