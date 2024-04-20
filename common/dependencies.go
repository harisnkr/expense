package common

import (
	log "log/slog"
	"os"

	"github.com/harisnkr/expense/config"
)

const (
	development = "development"
	staging     = "staging"
	production  = "production"
)

// SetDependencies sets dependencies based on environment
func SetDependencies() {
	log.Info("Setting dependencies first..")
	config.InitEnvVar()
	config.LoadECDSAKey()
	initValidators()
	initLogger()
}

func initLogger() {
	env := os.Getenv("MODE")
	if env != "" {
		log.Info("Initializing logger", "env detected", env)
	}

	switch env {
	case development, staging:
		log.SetLogLoggerLevel(log.LevelDebug)
	case production:
		log.SetLogLoggerLevel(log.LevelInfo)
	default:
		log.SetLogLoggerLevel(log.LevelDebug)
	}
}
