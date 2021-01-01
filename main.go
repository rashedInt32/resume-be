// Package main
package main

import (
	"resume/db"
	"resume/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

const port string = ":3001"

func main() {
	godotenv.Load(".env")

	db.Connect()

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New())

	routes.Setup(app)

	app.Listen(port)
}
