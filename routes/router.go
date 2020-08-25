package routes

import (
	controller "resume/controllers"

	"github.com/gofiber/fiber"
)

// Setup .
func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	user := api.Group("/user")
	user.Get("/check", func(c *fiber.Ctx) {
		c.Send("Working as expected")
	})

	user.Get("/auth/google/callback", controller.UserAuthentication)
}
