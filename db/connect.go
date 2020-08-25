// Package db to connect and create collection
package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// MgIns type will hold the client and bd
var MgIns mongoInstance

// Connect to the mongo db with mongo go client.
func Connect() {
	uri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Println("Couldn't connect to database.")
	}

	log.Println("Database connected")

	db := client.Database(dbName)

	MgIns = mongoInstance{
		Client: client,
		Db:     db,
	}
}
