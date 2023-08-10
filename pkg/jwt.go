package pkg

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const secret string = "your-secret-key"

type JWTInterface interface {
	Validate(tokenString string) (payload JWTClaims, statusCode int, err error)
	GenerateJWTToken(userID int64) (tokenString string, err error)
}

type JWTClaims struct {
	UserID int64
}

func NewJWT() (jwt *JWTClaims) {
	return &JWTClaims{}
}

func (j *JWTClaims) Validate(tokenString string) (payload JWTClaims, statusCode int, err error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return
	}

	// Check if the token is valid
	if !token.Valid {
		return payload, http.StatusForbidden, errors.New("Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return payload, http.StatusInternalServerError, errors.New("Internal Server Error")
	}
	expiration := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expiration) {
		return payload, http.StatusForbidden, errors.New("Token Has Expired")
	}
	payload.UserID = int64(claims["user_id"].(float64))

	return payload, 200, nil
}

func (j *JWTClaims) GenerateJWTToken(userID int64) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create the token using claims and a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error generating token:", err)
		return
	}

	return
}
