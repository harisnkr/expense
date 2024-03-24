package main

import (
	"log"

	"gorm.io/gorm"

	"github.com/harisnkr/expense/config"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/models"
)

var db *gorm.DB

func init() {
	config.InitEnvVar()
	db = data.InitDatabase()
}

func main() {
	// TODO: see if we can make this verbose
	err := db.AutoMigrate(models.Card{})
	if err != nil {
		log.Fatal(err, "migration failed")
		return
	}
}
