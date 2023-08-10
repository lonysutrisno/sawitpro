package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"sawitpro/mock"
	"sawitpro/pkg"
	"sawitpro/repository"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer_GetProfile(t *testing.T) {
	tests := []struct {
		name             string
		validateTokenErr error
		beforeTestJWT    func(*mock.MockJWTInterface)
		beforeTestRepo   func(*repository.MockRepositoryInterface)
		getUserErr       error
		expectedStatus   int
	}{
		{
			name:             "Valid Token and GetUserByUserID",
			validateTokenErr: nil,
			getUserErr:       nil,
			beforeTestJWT: func(mockJWT *mock.MockJWTInterface) {
				mockJWT.EXPECT().Validate(gomock.Any()).Return(pkg.JWTClaims{UserID: 1}, http.StatusOK, nil)
			},
			beforeTestRepo: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().GetUserByUserID(gomock.Any(), repository.GetUserByUserIDInput{UserID: 1}).Return(repository.GetUserByUserIDOutput{}, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:             "Invalid Token",
			validateTokenErr: errors.New("Unauthorized"),
			getUserErr:       nil,
			expectedStatus:   http.StatusForbidden,
			beforeTestJWT: func(mockJWT *mock.MockJWTInterface) {
				mockJWT.EXPECT().Validate(gomock.Any()).Return(pkg.JWTClaims{}, http.StatusForbidden, errors.New("Unauthorized"))
			},
			beforeTestRepo: nil,
		},
		{
			name:             "Valid Token but GetUserByUserID Error",
			validateTokenErr: nil,
			getUserErr:       errors.New("Database error"),
			expectedStatus:   http.StatusInternalServerError,
			beforeTestJWT: func(mockJWT *mock.MockJWTInterface) {
				mockJWT.EXPECT().Validate(gomock.Any()).Return(pkg.JWTClaims{UserID: 1}, http.StatusOK, nil)
			},
			beforeTestRepo: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().GetUserByUserID(gomock.Any(), repository.GetUserByUserIDInput{UserID: 1}).Return(repository.GetUserByUserIDOutput{}, errors.New("Error"))
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJWT := mock.NewMockJWTInterface(ctrl)
	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	s := &Server{
		JWT:        mockJWT,
		Repository: mockRepo,
	}

	e := echo.New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			req.Header.Set("Authorization", "Bearer valid_token") // Set your headers here

			ctx := e.NewContext(req, rec)
			if test.beforeTestJWT != nil {
				test.beforeTestJWT(mockJWT)
			}
			if test.beforeTestRepo != nil {
				test.beforeTestRepo(mockRepo)
			}
			s.GetProfile(ctx)

			assert.Equal(t, test.expectedStatus, ctx.Response().Status)
		})
	}
}

// func TestServer_PutProfile(t *testing.T) {
// 	tests := []struct {
// 		name            string
// 		jsonBody        *generated.PutProfileJSONRequestBody
// 		expectedStatus  int
// 		expectedMessage string
// 		jwtErr          error
// 		updateUserErr   error
// 	}{
// 		{
// 			name: "Valid Request",
// 			jsonBody: &generated.PutProfileJSONRequestBody{
// 				FullName:    "John Doe",
// 				PhoneNumber: "1234567890",
// 			},
// 			expectedStatus:  http.StatusOK,
// 			expectedMessage: "success",
// 			jwtErr:          nil,
// 			updateUserErr:   nil,
// 		},
// 		{
// 			name: "Invalid JSON Body",
// 			jsonBody: &generated.PutProfileJSONRequestBody{
// 				FullName:    nil, // Invalid JSON body
// 				PhoneNumber: "1234567890",
// 			},
// 			expectedStatus:  http.StatusBadRequest,
// 			expectedMessage: "Invalid full name: unexpected end of JSON input",
// 			jwtErr:          nil,
// 			updateUserErr:   nil,
// 		},
// 		{
// 			name: "Invalid Full Name",
// 			jsonBody: &generated.PutProfileJSONRequestBody{
// 				FullName:    "A", // Invalid full name
// 				PhoneNumber: "1234567890",
// 			},
// 			expectedStatus:  http.StatusBadRequest,
// 			expectedMessage: "Invalid full name: must be at least 3 characters",
// 			jwtErr:          nil,
// 			updateUserErr:   nil,
// 		},
// 		{
// 			name: "Invalid Phone Number",
// 			jsonBody: &generated.PutProfileJSONRequestBody{
// 				FullName:    "John Doe",
// 				PhoneNumber: "abc", // Invalid phone number
// 			},
// 			expectedStatus:  http.StatusBadRequest,
// 			expectedMessage: "Invalid phone number: must be a valid phone number",
// 			jwtErr:          nil,
// 			updateUserErr:   nil,
// 		},
// 		{
// 			name: "JWT Validation Error",
// 			jsonBody: &generated.PutProfileJSONRequestBody{
// 				FullName:    "John Doe",
// 				PhoneNumber: "1234567890",
// 			},
// 			expectedStatus:  http.StatusForbidden,
// 			expectedMessage: "Unauthorized",
// 			jwtErr:          errors.New("Unauthorized"),
// 			updateUserErr:   nil,
// 		},
// 		{
// 			name: "Update User Error (Phone Number Conflict)",
// 			jsonBody: generated.PutProfileJSONRequestBody{
// 				FullName:    "John Doe",
// 				PhoneNumber: "1234567890",
// 			},
// 			expectedStatus:  http.StatusConflict,
// 			expectedMessage: "phone number already exists",
// 			jwtErr:          nil,
// 			updateUserErr:   errors.New("phone number already exists"),
// 		},
// 		{
// 			name: "Update User Error (Internal Server Error)",
// 			jsonBody: generated.PutProfileJSONRequestBody{
// 				FullName:    "John Doe",
// 				PhoneNumber: "1234567890",
// 			},
// 			expectedStatus:  http.StatusInternalServerError,
// 			expectedMessage: "Database error",
// 			jwtErr:          nil,
// 			updateUserErr:   errors.New("Database error"),
// 		},
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockJWT := mock.NewMockJWTInterface(ctrl)
// 	mockRepo := repository.NewMockRepositoryInterface(ctrl)

// 	s := &Server{
// 		JWT:        mockJWT,
// 		Repository: mockRepo,
// 	}

// 	e := echo.New()

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			req := httptest.NewRequest(http.MethodPut, "/", nil)
// 			rec := httptest.NewRecorder()

// 			body, _ := json.Marshal(test.jsonBody)
// 			req.Header.Set("Content-Type", "application/json")
// 			req.Body = ioutil.NopCloser(bytes.NewReader(body))

// 			c := e.NewContext(req, rec)
// 			// ctx := context.Background()

// 			// Set up mock expectations and behavior
// 			mockJWT.EXPECT().Validate(gomock.Any()).Return(pkg.JWTClaims{}, test.expectedStatus, test.jwtErr)

// 			if test.expectedStatus != http.StatusOK {
// 				// errResponse := generated.ErrorResponse{Message: test.expectedMessage}
// 				mockRepo.EXPECT().UpdateUserByUserID(gomock.Any(), gomock.Any()).Return(test.updateUserErr)
// 				s.PutProfile(c)
// 				assert.JSONEq(t, `{"message":"`+test.expectedMessage+`"}`, rec.Body.String())
// 				assert.Equal(t, test.expectedStatus, rec.Code)
// 			} else {
// 				// UpdateUserByUserID should not be called for successful scenarios
// 				mockRepo.EXPECT().UpdateUserByUserID(gomock.Any(), gomock.Any()).Times(0)
// 				s.PutProfile(c)
// 				assert.Empty(t, rec.Body.String())
// 				assert.Equal(t, test.expectedStatus, rec.Code)
// 			}

// 		})
// 	}
// }
