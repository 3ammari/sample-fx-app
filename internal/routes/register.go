package routes

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/3ammari/sample-fx-app/internal/env"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/3ammari/sample-fx-app/internal/handler/hello"
)

// Params is the input parameter struct for the module that contains its
// dependencies.
type Params struct {
	fx.In
	DB        *sql.DB
	Logger    *zap.Logger
	Lifecycle fx.Lifecycle
	Handler   hello.Handler
	Config    env.Config
}

// Register registers the routes for the server and starts the server on app
// start.
func Register(p Params) {
	router := http.NewServeMux()
	router.HandleFunc("/", p.Handler.Hello)
	server := http.Server{
		Addr:    ":" + p.Config.Port,
		Handler: router,
	}

	p.Lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				p.Logger.Info("Starting server.")
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				p.Logger.Info("Shutting down server.")
				return server.Shutdown(ctx)
			},
		},
	)
}
