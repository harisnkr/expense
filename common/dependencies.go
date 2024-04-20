package common

import (
	log "log/slog"
	"os"

	"github.com/harisnkr/expense/config"
)

const (
	// Development is the local/test environment
	Development = "development"
	// Staging is the UAT/pre-production environment
	Staging = "staging"
	// Production is the live environment
	Production = "production"
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
	case Development, Staging:
		log.SetLogLoggerLevel(log.LevelDebug)
	case Production:
		log.SetLogLoggerLevel(log.LevelInfo)
	default:
		log.SetLogLoggerLevel(log.LevelDebug)
	}
}

// GetMode gets the current environment mode that this service is running on
func GetMode() string {
	if env := os.Getenv("MODE"); env != "" {
		return env
	}
	log.Warn("undefined env, defaulting to development")
	return Development
}
