// Package utils provides utility for app
package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generate jwt token.
func GenerateToken(email string) (t string, err error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return t, err
}
