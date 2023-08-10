package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sawitpro/generated"
	"sawitpro/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) PostLogin(ctx echo.Context) error {
	jsonBody := generated.PostLoginJSONRequestBody{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	err = validateBodyLogin(jsonBody)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	userpassword, err := s.Repository.GetUserPassword(context.Background(), repository.GetUserPasswordInput{PhoneNumber: *jsonBody.PhoneNumber})
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}

	err = compareSaltedPasswords(userpassword.Password, userpassword.PasswordSalt, *jsonBody.Password)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusForbidden, generated.ErrorResponse{Message: err.Error()})
	}
	err = s.Repository.UpdateSuccessLoginByUserID(context.Background(), repository.UpdateSuccessLoginByUserIDInput{PhoneNumber: *jsonBody.PhoneNumber})
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	token, err := s.JWT.GenerateJWTToken(userpassword.UserID)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func validateBodyLogin(param generated.PostLoginJSONRequestBody) (err error) {
	if param.PhoneNumber == nil {
		return fmt.Errorf("Required phone_number")
	}
	err = validatePhoneNumber(*param.PhoneNumber)
	if err != nil {
		return fmt.Errorf("Invalid phone_number:" + err.Error())
	}

	return
}

func compareSaltedPasswords(hashedPassword, salt, plainPassword string) error {
	// Re-create the salted password for comparison
	saltedPassword := plainPassword + salt

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(saltedPassword))
	if err != nil {
		return fmt.Errorf("password mismatch")
	}
	return nil
}
