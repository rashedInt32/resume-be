// Package db to connect and create collection
package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = "mongodb://localhost:27017"

// Connect to the mongo db with mongo go client.
func Connect() {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Println("Couldn't connect to database.")
	}

	log.Println("Database connected")

	db := client.Database("resume")

	db.Collection("user")

}
