// Package main provides run the app
package main

import (
	"resume/db"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

const port int = 8080

func main() {
	db.Connect()
	app := fiber.New()

	app.Use(middleware.Logger("${method} ${path} ${latecy}"))

	app.Get("/user", func(c *fiber.Ctx) {
		c.Send("App running on port 9000")
	})

	app.Get("/user/post", func(c *fiber.Ctx) {
		c.Send("Post api")
	})

	app.Listen(port)
}
