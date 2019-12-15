package loggerfx

import (
	"github.com/3ammari/sample-fx-app/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module is the loggerfx module that can be passed into an Fx app.
var Module = fx.Provide(New)

//Params New Input
type Params struct {
	fx.In
	Config env.Config
}

// New constructs a new logger.
func New(p Params) (logger *zap.Logger, err error) {
	switch p.Config.Stage {
	case env.Production:
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	logger.Info("Logger in Stage " + string(p.Config.Stage) + " Mode")
	return logger, nil
}
