// Package middleware provides to protect routes
package middleware

import (
	"os"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

// Protected route.
func Protected() func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}
