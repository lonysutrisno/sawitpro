package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sawitpro/generated"
	"sawitpro/pkg"
	"sawitpro/repository"
	"strings"

	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
func (s *Server) PostRegister(ctx echo.Context) error {

	jsonBody := generated.PostRegisterJSONBody{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}

	err = validateBody(jsonBody)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	salt := pkg.GenerateRandomSalt()
	hashedPassword, err := pkg.HashPassword(*jsonBody.Password, salt)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	id, err := s.Repository.InsertUser(context.Background(), repository.UserInput{
		Fullname:     *jsonBody.FullName,
		PhoneNumber:  *jsonBody.PhoneNumber,
		Password:     hashedPassword,
		PasswordSalt: salt})
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, id)
}

func validateBody(param generated.PostRegisterJSONBody) (err error) {
	if param.FullName == nil || param.Password == nil || param.PhoneNumber == nil {
		return fmt.Errorf("Required phone_number, full_name, password")
	}
	if err := validatePhoneNumber(*param.PhoneNumber); err != nil {
		return fmt.Errorf("Invalid phone number:" + err.Error())
	}

	if err := validateFullName(*param.FullName); err != nil {
		return fmt.Errorf("Invalid full name:" + err.Error())
	}

	if err := validatePassword(*param.Password); err != nil {
		return fmt.Errorf("Invalid password:" + err.Error())
	}
	return
}

func validatePhoneNumber(phoneNumber string) error {
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return fmt.Errorf("phone number must be between 10 and 13 characters")
	}

	if !strings.HasPrefix(phoneNumber, "+62") {
		return fmt.Errorf("phone number must start with +62")
	}

	return nil
}

func validateFullName(fullName string) error {
	if len(fullName) < 3 || len(fullName) > 60 {
		return fmt.Errorf("full name must be between 3 and 60 characters")
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 || len(password) > 64 {
		return fmt.Errorf("password must be between 6 and 64 characters")
	}

	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialRegex := regexp.MustCompile(`[^A-Za-z0-9]`)

	if !uppercaseRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least 1 uppercase letter")
	}

	if !lowercaseRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least 1 lowercase letter")
	}

	if !digitRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least 1 digit")
	}

	if !specialRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least 1 special character")
	}

	return nil
}
