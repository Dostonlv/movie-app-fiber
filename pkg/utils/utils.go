package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
