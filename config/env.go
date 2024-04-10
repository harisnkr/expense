package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var ECDSAKey *ecdsa.PrivateKey

// InitEnvVar initialises environment variables declared in ../.env
func InitEnvVar() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// LoadECDSAKey loads ECDSA private key from .env, or generates a new one for dev
func LoadECDSAKey() {
	keyFromEnv := os.Getenv("ECDSA_PRIVATE_KEY")
	if keyFromEnv != "" {
		if err := setECDSAKeyFromEnv(keyFromEnv); err != nil {
			log.Fatal("error loading ECDSA key from .env file: ", err)
		}
		return
	}
	generateRandomECDSAKey()

}

func setECDSAKeyFromEnv(keyFromEnv string) error {
	keyBytes, err := base64.URLEncoding.DecodeString(keyFromEnv)
	if err != nil {
		log.Fatal("error decoding base64 string", err)
	}

	privateKey, err := x509.ParseECPrivateKey(keyBytes)
	if err != nil {
		return errors.New("error parsing ECDSA private key")
	}

	ECDSAKey = privateKey
	return nil
}

func generateRandomECDSAKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal("Error generating ECDSA private key: ", err)
		return
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatal("Error marshaling ECDSA private key", err)
		return
	}
	ECDSAKey = privateKey
	privateKeyBase64 := base64.URLEncoding.EncodeToString(privateKeyBytes)

	log.Warn("ECDSA_PRIVATE_KEY environment variable not set, generated random key for testing: " + privateKeyBase64)
}
