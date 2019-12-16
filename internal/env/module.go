package env

import (
	"os"

	"go.uber.org/fx"
)

//Module Environment Store
var Module = fx.Provide(New)

//Stage stage enum
type Stage string

const (
	Dev        = "Dev"
	Staging    = "Staging"
	Production = "Production"
)

//Config is environment vars store
type Config struct {
	DatabaseURL string
	Stage       Stage
	Port        string
}

//New instantiate an new Default
func New() Config {
	config := Config{
		DatabaseURL: os.Getenv("DB_URL"),
		Port:        os.Getenv("PORT"),
	}

	switch os.Getenv("Stage") {
	case "Prod":
		config.Stage = Production
	case "Staging":
		config.Stage = Staging
	default:
		config.Stage = Dev
	}

	return config
}
