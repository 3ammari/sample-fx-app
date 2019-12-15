package main

import (
	"github.com/3ammari/sample-fx-app/internal/env"
	"go.uber.org/fx"

	"github.com/3ammari/sample-fx-app/internal/db"
	"github.com/3ammari/sample-fx-app/internal/handler"
	"github.com/3ammari/sample-fx-app/internal/loggerfx"
	"github.com/3ammari/sample-fx-app/internal/routes"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		handler.Module,
		loggerfx.Module,
		db.Module,
		env.Module,
		fx.Invoke(routes.Register),
	)
}
