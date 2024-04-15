package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"errors"
	log "log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/harisnkr/expense/common"
)

var (
	// ECDSAKey is the loaded/generated private key for token generation
	ECDSAKey *ecdsa.PrivateKey

	// SessionTokenTTLInHours is the loaded/configured session token TTL
	SessionTokenTTLInHours time.Duration

	defaultSessionTokenTTL = time.Hour * 24 * 30 * 3
)

// InitEnvVar initialises environment variables declared in ../.env
func InitEnvVar() {
	if err := godotenv.Load(); err != nil {
		log.Error("Error loading .env file")
	}
	setTokenTTLConfig()
}

func setTokenTTLConfig() {
	durationCfg, err := time.ParseDuration(os.Getenv("TOKEN_TTL"))
	if err != nil {
		log.Error(common.Error, err, "error parsing TOKEN_TTL from .env")
		log.Info("Setting session token TTL to default value",
			"SessionTokenTTLInHours", defaultSessionTokenTTL)
		SessionTokenTTLInHours = defaultSessionTokenTTL // 3 months
		return
	}
	// if valid TTL config found
	SessionTokenTTLInHours = durationCfg * time.Hour
	log.Info("Setting TOKEN_TTL in hours", "SessionTokenTTLInHours", SessionTokenTTLInHours)
}

// LoadECDSAKey loads ECDSA private key from .env, or generates a new one for dev
func LoadECDSAKey() {
	keyFromEnv := os.Getenv("ECDSA_PRIVATE_KEY")
	if keyFromEnv != "" {
		if err := setECDSAKeyFromEnv(keyFromEnv); err != nil {
			log.Error(common.Error, err, "error loading ECDSA key from .env file")
		}
		return
	}
	generateRandomECDSAKey()
}

func setECDSAKeyFromEnv(keyFromEnv string) error {
	keyBytes, err := base64.URLEncoding.DecodeString(keyFromEnv)
	if err != nil {
		return errors.New("error decoding base64 string")
	}

	privateKey, err := x509.ParseECPrivateKey(keyBytes)
	if err != nil {
		return errors.New("error parsing ECDSA private key")
	}

	ECDSAKey = privateKey
	log.Info("Successfully loaded ECDSA key from .env file")
	return nil
}

func generateRandomECDSAKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Error(common.Error, err, "error generating ECDSA private key")
		return
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Error(common.Error, err, "error marshaling ECDSA private key")
		return
	}
	ECDSAKey = privateKey
	privateKeyBase64 := base64.URLEncoding.EncodeToString(privateKeyBytes)

	log.Warn("ECDSA_PRIVATE_KEY environment variable not set, generated random key for testing: " + privateKeyBase64)
}
