package routes

import (
	"resume/controller"

	"github.com/gofiber/fiber"
)

// Setup .
func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	user := api.Group("/user")
	user.Get("/auth", controller.UserAuth)
	user.Post("/signup", controller.UserSignup)
}
