package routes

import (
	"resume/controller"
	"resume/middleware"

	"github.com/gofiber/fiber/v2"
)

// Setup .
func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	user := api.Group("/user")
	user.Post("/auth", controller.Auth)
	user.Post("/signup", controller.Signup)
	user.Post("/resume", middleware.Protected(), controller.Resume)
	user.Get("/resume/:id", middleware.Protected(), controller.GetResume)
	user.Get("/me", middleware.Protected(), controller.Me)
}
