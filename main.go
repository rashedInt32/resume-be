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

const port int = 3001

func main() {
	godotenv.Load(".env")

	db.Connect()

	app := fiber.New()

	app.Use(cors.New())

	app.Use(middleware.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	routes.Setup(app)

	app.Listen(port)
}
