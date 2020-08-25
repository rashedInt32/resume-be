// Package main
package main

import (
	"log"
	"os"
	"resume/db"
	"resume/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

const port int = 8080

func main() {
	godotenv.Load(".env")

	engine := html.New("./views", ".html")

	db.Connect()

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback", "email"),
	)

	app.Get("/auth/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Send(user)

	})

	app.Use(middleware.Logger("${method} ${path} ${latecy}"))

	routes.Setup(app)

	app.Listen(port)
}
