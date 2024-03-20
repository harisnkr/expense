package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// InitDatabase initalises the Postgres database
func InitDatabase() {
	dsn := os.Getenv("DSN")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
}
