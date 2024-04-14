package data

import (
	"context"
	log "log/slog"
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
		log.Error("failed to connect to mongodb", err)
		panic(err)
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Error("Failed to connect to database")
		panic(err)
	}
	log.Info("Connected to MongoDB!")

	return client, &Collections{
		Cards: client.Database(databaseName).Collection(cardsCollection),
		Users: client.Database(databaseName).Collection(usersCollection),
	}
}
