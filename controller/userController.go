// Package controller provides
package controller

import (
	"context"
	"log"
	"resume/db"
	"resume/model"
	"resume/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

// GetUserCollection .
func GetUserCollection() *mongo.Collection {
	return db.MgIns.DB.Collection("user")
}

// Auth : Post login.
func Auth(c *fiber.Ctx) error {
	userCollection = GetUserCollection()
	ctx := context.Background()

	var user model.Auth
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	var result model.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "No User found on that email",
		})
	}

	isPasswordMatch := utils.CheckHashPassword(user.Password, result.Password)
	if !isPasswordMatch {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "Password not match",
		})
	}

	token, _ := utils.GenerateToken(result.Email)

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}

// Signup : post registration user.
func Signup(c *fiber.Ctx) error {
	userCollection = GetUserCollection()
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

	count, _ := userCollection.CountDocuments(ctx, filter)

	if count >= 1 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"msg":    "Username or email already taken, try different one.",
		})
	}

	pass, _ := utils.HashPassword(user.Password)
	user.Password = pass
	userCollection.InsertOne(ctx, user)

	token, _ := utils.GenerateToken(user.Email)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"token":  token,
	})
}

// Me : Get current user data.
func Me(c *fiber.Ctx) error {
	userCollection = GetUserCollection()
	ctx := context.Background()

	email := utils.ValidateToken(c)

	var result model.User

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"data": "Something went wrong",
		})
	}

	result.Password = "ha ha ha ho ho ha ha"
	return c.Status(200).JSON(fiber.Map{
		"user": result,
	})
}
