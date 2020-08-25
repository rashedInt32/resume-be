// Package controller provides
package controller

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

// UserAuthentication google 0auth.
func UserAuthentication(c *fiber.Ctx) {

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://127.0.0.1:8080/api/v1/user/auth/google/callback"),
	)

	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}

	c.Send(user)

}
