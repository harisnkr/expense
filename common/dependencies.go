package common

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/harisnkr/expense/config"
)

const (
	development = "development"
	staging     = "staging"
	production  = "production"
)

// SetDependencies sets dependencies based on environment
func SetDependencies() {
	config.InitEnvVar()
	config.LoadECDSAKey()
	setLogLevel()
	initValidators()
}

func setLogLevel() {
	env := os.Getenv("MODE")
	log.Info("detected environment mode: ", env)
	switch env {
	case development, staging:
		log.SetLevel(log.DebugLevel)
	case production:
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}
