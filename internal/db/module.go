package db

import (
	"context"
	"database/sql"

	"go.uber.org/zap"

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
	Logger *zap.Logger
	LC     fx.Lifecycle
}

//New creates a db connection
func New(p Params) (*sql.DB, error) {
	db, err := sql.Open("mysql", p.Config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			p.Logger.Sugar().Info("Ping DB")
			return db.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Sugar().Info("Closing DB connection")
			return db.Close()
		},
	})
	return db, nil
}
