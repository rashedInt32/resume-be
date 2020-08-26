// Package main
package main

import (
	"resume/db"
	"resume/routes"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"
)

const port int = 8080

func main() {
	godotenv.Load(".env")

	db.Connect()

	app := fiber.New()

	app.Use(cors.New())

	app.Use(middleware.Logger("${method} ${path} ${latecy}"))

	routes.Setup(app)

	app.Listen(port)
}
