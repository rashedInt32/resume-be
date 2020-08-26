// Package controller provides
package controller

import (
	"context"
	"encoding/json"
	"log"
	"resume/db"
	model "resume/models"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// UserAuth google 0auth.
func UserAuth(c *fiber.Ctx) {

	c.Send("user auth")

}

// UserSignup sign up user.
func UserSignup(c *fiber.Ctx) {
	collection = db.MgIns.Db.Collection("user")
	ctx := context.Background()
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	count, _ := collection.CountDocuments(ctx, bson.M{"email": user.Email})

	if count >= 1 {
		c.Send("User exists")
		return
	}

	collection.InsertOne(ctx, user)

	json, _ := json.Marshal(user)
	c.Send(json)
}
