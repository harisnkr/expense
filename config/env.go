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
)

var (
	// ECDSAKey is the loaded/generated private key for token generation
	ECDSAKey *ecdsa.PrivateKey

	// SessionTokenTTLInHours is the loaded/configured session token TTL
	SessionTokenTTLInHours time.Duration
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
		log.Error("Error parsing TOKEN_TTL")
		SessionTokenTTLInHours = time.Hour * 24 * 30 * 3 // 3 months
		return
	}
	// if valid TTL config found
	durationCfg = durationCfg * time.Hour
	if SessionTokenTTLInHours == 0 {
		SessionTokenTTLInHours = time.Hour * 24 * 30
	}
}

// LoadECDSAKey loads ECDSA private key from .env, or generates a new one for dev
func LoadECDSAKey() {
	keyFromEnv := os.Getenv("ECDSA_PRIVATE_KEY")
	if keyFromEnv != "" {
		if err := setECDSAKeyFromEnv(keyFromEnv); err != nil {
			log.Error("error loading ECDSA key from .env file: ", err)
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
		log.Error("Error generating ECDSA private key: ", err)
		return
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Error("Error marshaling ECDSA private key", err)
		return
	}
	ECDSAKey = privateKey
	privateKeyBase64 := base64.URLEncoding.EncodeToString(privateKeyBytes)

	log.Warn("ECDSA_PRIVATE_KEY environment variable not set, generated random key for testing: " + privateKeyBase64)
}
