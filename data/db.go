package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabase initalises the Postgres database. TODO: change to MySQL
func InitDatabase() *gorm.DB {
	dsn := os.Getenv("DSN")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	return database
}
