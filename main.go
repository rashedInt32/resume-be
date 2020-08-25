// Package main
package main

import (
	"encoding/json"
	"log"
	"os"
	"resume/db"
	"resume/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/session"
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

	// create session handler
	sessions := session.New()

	goth_fiber.Session = sessions

	log.Println(sessions)

	app.Get("/auth/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Fatal(err)
		}

		json, _ := json.Marshal(user)

		ctx.Send(json)

	})

	app.Use(middleware.Logger("${method} ${path} ${latecy}"))

	routes.Setup(app)

	app.Listen(port)
}
