package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ... TODO: dependency injection in main.go for routes
var DB *gorm.DB

// InitDatabase initalises the Postgres database. TODO: change to MySQL
func InitDatabase() {
	dsn := os.Getenv("DSN")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
}
