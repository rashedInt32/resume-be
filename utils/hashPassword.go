// Package utils provides utility for the app
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword bcrypt the password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckHashPassword compare password match.
func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
