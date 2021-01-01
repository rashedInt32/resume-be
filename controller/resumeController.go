package controller

import (
	"context"
	"log"
	"resume/db"
	"resume/model"
	"resume/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var resumeCollection *mongo.Collection

// getResumeCollection .
func getResumeCollection() *mongo.Collection {
	return db.MgIns.DB.Collection("resume")
}

// GetResume to get specific resume.
func GetResume(c *fiber.Ctx) error {
	resumeCollection := getResumeCollection()
	ctx := context.Background()

	id := c.Params("id")
	resumeID, err := primitive.ObjectIDFromHex(id)

	var resume model.Resume
	err = resumeCollection.FindOne(ctx, bson.M{"_id": resumeID}).Decode(&resume)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"resume": resume})
}

// Resume : Post resume to associate users
func Resume(c *fiber.Ctx) error {
	resumeCollection := getResumeCollection()
	userCollection := GetUserCollection()
	ctx := context.Background()

	var resumedata model.Resume
	if err := c.BodyParser(&resumedata); err != nil {
		log.Fatal(err)
	}

	resume, err := resumeCollection.InsertOne(ctx, resumedata)
	if err != nil {
		log.Fatal(err)
	}

	email := utils.ValidateToken(c)

	result := userCollection.FindOneAndUpdate(ctx, bson.M{
		"email": email,
	},
		bson.M{
			"$set": bson.M{"resume": resume.InsertedID},
		})

	if result.Err() != nil {
		log.Panicln(result.Err())
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"msg":    "Resume successfully added",
	})
}
