package global

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB holds database connection
var DB mongo.Database

// ConnectToDB create an mongo database connection
func ConnectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal("Error in connecting db", err.Error())
	}
	DB = *client.Database(dbname)
}

// NewDBContext returns new context
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*performance/100)
}

// Testing database
func ConnectToTestDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal("Error in connecting db", err.Error())
	}
	DB = *client.Database(dbname + "_test")
}
