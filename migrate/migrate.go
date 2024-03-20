package main

import (
	"log"

	"github.com/harisnkr/expense/config"
	"github.com/harisnkr/expense/data"
	"github.com/harisnkr/expense/models"
)

func init() {
	config.InitEnvVar()
	data.InitDatabase()
}

func main() {
	// TODO: see if we can make this verbose
	err := data.DB.AutoMigrate(models.Card{})
	if err != nil {
		log.Fatal(err, "migration failed")
		return
	}
}
