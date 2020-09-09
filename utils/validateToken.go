package utils

import "github.com/dgrijalva/jwt-go"

// ValidateToken token to valid.
func ValidateToken(t *jwt.Token, user string) string {
	claims := t.Claims.(jwt.MapClaims)
	email := string(claims["email"].(string))

	return email
}
