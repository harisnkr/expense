package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName = "expense"

	cardsCollection = "cards"
	usersCollection = "users"
)

// Collections ...
type Collections struct {
	Cards *mongo.Collection
	Users *mongo.Collection
}

// InitDatabase inits MongoDB and its collections
func InitDatabase(ctx context.Context) (*mongo.Client, *Collections) {
	dsn := os.Getenv("DSN")
	clientOptions := options.Client().ApplyURI(dsn)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	return client, &Collections{
		Cards: client.Database(databaseName).Collection(cardsCollection),
		Users: client.Database(databaseName).Collection(usersCollection),
	}
}
