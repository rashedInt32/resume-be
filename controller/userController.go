// Package controller provides
package controller

import (
	"context"
	"log"
	"resume/db"
	"resume/model"
	"resume/utils"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Auth : Post login.
func Auth(c *fiber.Ctx) {
	collection = db.MgIns.Db.Collection("user")
	ctx := context.Background()

	var user model.Auth
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	var result model.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&result)
	if err != nil {
		c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "No User found on that email",
		})
		return
	}

	isPasswordMatch := utils.CheckHashPassword(user.Password, result.Password)
	if !isPasswordMatch {
		c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "Password not match",
		})
		return
	}

	token, _ := utils.GenerateToken(result.Email)

	c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}

// Signup : post registration user.
func Signup(c *fiber.Ctx) {
	collection = db.MgIns.Db.Collection("user")
	ctx := context.Background()
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"email": user.Email},
			bson.M{"username": user.Username},
		},
	}

	count, _ := collection.CountDocuments(ctx, filter)

	if count >= 1 {
		c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "Username or email already taken, try different one.",
		})
		return
	}

	pass, _ := utils.HashPassword(user.Password)
	user.Password = pass
	collection.InsertOne(ctx, user)

	token, _ := utils.GenerateToken(user.Email)

	c.Status(200).JSON(fiber.Map{
		"status": "success",
		"token":  token,
	})
}

// Resume : Post resume to associate users
func Resume(c *fiber.Ctx) {
	collection = db.MgIns.Db.Collection("user")
	ctx := context.Background()

	var resumedata model.Resume
	if err := c.BodyParser(&resumedata); err != nil {
		log.Fatal(err)
	}

	email := utils.ValidateToken(c)

	var result model.User
	err := collection.FindOneAndUpdate(ctx, bson.M{
		"email": email,
	},
		bson.M{
			"$set": bson.M{"resume": resumedata},
		}).Decode(&result)
	if err != nil {
		c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "No User found on that email",
		})
		return
	}

	c.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

// Me get current user data.
func Me(c *fiber.Ctx) {
	collection = db.MgIns.Db.Collection("user")
	ctx := context.Background()

	email := utils.ValidateToken(c)

	var result model.User

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)

	if err != nil {
		c.Status(404).JSON(fiber.Map{
			"data": "Something went wrong",
		})
		return
	}

	result.Password = "ha ha ha ho ho ha ha"
	c.Status(200).JSON(result)
}
