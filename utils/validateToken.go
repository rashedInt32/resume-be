package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

// ValidateToken token to valid.
func ValidateToken(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	return email
}
