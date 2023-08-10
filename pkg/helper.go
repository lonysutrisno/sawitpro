package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ExtractToken(ctx echo.Context) (token string) {
	authHeader := ctx.Request().Header.Get("Authorization")
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func GenerateRandomSalt() string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return base64.URLEncoding.EncodeToString(salt)
}

func HashPassword(password, salt string) (string, error) {
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)

	hashedPassword, err := bcrypt.GenerateFromPassword(append(passwordBytes, saltBytes...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
