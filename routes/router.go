package routes

import (
	"os"
	"resume/controller"
	"resume/middleware"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

// Setup .
func Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	user := api.Group("/user")
	user.Post("/auth", controller.UserAuth)
	user.Post("/signup", controller.UserSignup)
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	user.Post("/resume", middleware.Protected(), controller.UserResume)
}
