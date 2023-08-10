package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sawitpro/generated"
	"sawitpro/pkg"
	"sawitpro/repository"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetProfile(ctx echo.Context) error {
	restoken, statusCode, err := s.JWT.Validate(pkg.ExtractToken(ctx))
	if err != nil {
		log.Println(err)
		return ctx.JSON(statusCode, generated.ErrorResponse{Message: err.Error()})
	}
	res, err := s.Repository.GetUserByUserID(context.Background(), repository.GetUserByUserIDInput{UserID: restoken.UserID})
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(statusCode, res)
}

func (s *Server) PutProfile(ctx echo.Context) error {
	jsonBody := generated.PutProfileJSONRequestBody{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}
	param, err := validateBodyUpdateUser(jsonBody)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
	}

	restoken, statusCode, err := s.JWT.Validate(pkg.ExtractToken(ctx))
	if err != nil {
		log.Println(err)
		return ctx.JSON(statusCode, generated.ErrorResponse{Message: err.Error()})
	}
	param.UserID = restoken.UserID

	err = s.Repository.UpdateUserByUserID(context.Background(), param)
	if err != nil {
		if err.Error() == "phone number already exists" {
			return ctx.JSON(http.StatusConflict, generated.ErrorResponse{Message: err.Error()})
		}
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, generated.ErrorResponse{Message: "success"})
}

func validateBodyUpdateUser(param generated.PutProfileJSONRequestBody) (nonNilParam repository.UpdateUserByUserIDInput, err error) {
	if param.FullName != nil {
		nonNilParam.Fullname = *param.FullName
		err = validateFullName(*param.FullName)
		if err != nil {
			return nonNilParam, fmt.Errorf("Invalid full_name:" + err.Error())
		}
	}
	if param.PhoneNumber != nil {
		nonNilParam.PhoneNumber = *param.PhoneNumber
		err = validatePhoneNumber(*param.PhoneNumber)
		if err != nil {
			return nonNilParam, fmt.Errorf("Invalid phone_number:" + err.Error())
		}
	}
	return
}
