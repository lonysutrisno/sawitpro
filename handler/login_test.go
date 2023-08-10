package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sawitpro/generated"
	"sawitpro/mock"
	"sawitpro/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer_PostLogin(t *testing.T) {
	tests := []struct {
		name             string
		validateTokenErr error
		phone            string
		password         string
		beforeTestJWT    func(*mock.MockJWTInterface)
		beforeTestRepo   func(*repository.MockRepositoryInterface)
		getUserErr       error
		expectedStatus   int
	}{
		{
			name:       "Success Login",
			getUserErr: nil,
			phone:      "+62123459012",
			password:   "Pass1234!",
			beforeTestRepo: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().GetUserPassword(gomock.Any(), repository.GetUserPasswordInput{PhoneNumber: "+62123459012"}).Return(repository.GetUserPasswordOutput{
					Password:     "$2a$10$6EC8qw9D0qWjqvdiOHZaIO4zGzLtqpAOC5i1gOA1PCfcS08y3jC5i",
					PasswordSalt: "RPXk3rahIhvP7Tf-UhHAYw==",
					UserID:       1,
				}, nil)
				mockRepo.EXPECT().UpdateSuccessLoginByUserID(gomock.Any(), repository.UpdateSuccessLoginByUserIDInput{PhoneNumber: "+62123459012"}).Return(nil)
			},
			beforeTestJWT: func(mockJWT *mock.MockJWTInterface) {
				mockJWT.EXPECT().GenerateJWTToken(int64(1)).Return("sometoken", nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:       "False Password Login",
			getUserErr: nil,
			phone:      "+62123459012",
			password:   "Pass1234",
			beforeTestRepo: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().GetUserPassword(gomock.Any(), repository.GetUserPasswordInput{PhoneNumber: "+62123459012"}).Return(repository.GetUserPasswordOutput{
					Password:     "$2a$10$6EC8qw9D0qWjqvdiOHZaIO4zGzLtqpAOC5i1gOA1PCfcS08y3jC5i",
					PasswordSalt: "RPXk3rahIhvP7Tf-UhHAYw==",
					UserID:       1,
				}, nil)
			},

			expectedStatus: http.StatusForbidden,
		},
		{
			name:       "Failed generate token Login",
			getUserErr: nil,
			phone:      "+62123459012",
			password:   "Pass1234!",
			beforeTestRepo: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().GetUserPassword(gomock.Any(), repository.GetUserPasswordInput{PhoneNumber: "+62123459012"}).Return(repository.GetUserPasswordOutput{
					Password:     "$2a$10$6EC8qw9D0qWjqvdiOHZaIO4zGzLtqpAOC5i1gOA1PCfcS08y3jC5i",
					PasswordSalt: "RPXk3rahIhvP7Tf-UhHAYw==",
					UserID:       1,
				}, nil)
				mockRepo.EXPECT().UpdateSuccessLoginByUserID(gomock.Any(), repository.UpdateSuccessLoginByUserIDInput{PhoneNumber: "+62123459012"}).Return(nil)
			},
			beforeTestJWT: func(mockJWT *mock.MockJWTInterface) {
				mockJWT.EXPECT().GenerateJWTToken(int64(1)).Return("", errors.New("some error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:       "Failed Get Password",
			getUserErr: nil,
			phone:      "+62123459012",
			password:   "Pass1234!",
			beforeTestRepo: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().GetUserPassword(gomock.Any(), repository.GetUserPasswordInput{PhoneNumber: "+62123459012"}).Return(repository.GetUserPasswordOutput{}, errors.New("some error"))
			},

			expectedStatus: http.StatusInternalServerError,
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
			param := generated.PostLoginJSONRequestBody{
				PhoneNumber: &test.phone,
				Password:    &test.password,
			}
			req.Header.Set("Authorization", "Bearer valid_token") // Set your headers here
			body, _ := json.Marshal(param)
			req.Header.Set("Content-Type", "application/json")
			req.Body = ioutil.NopCloser(bytes.NewReader(body))

			ctx := e.NewContext(req, rec)

			if test.beforeTestJWT != nil {
				test.beforeTestJWT(mockJWT)
			}
			if test.beforeTestRepo != nil {
				test.beforeTestRepo(mockRepo)
			}
			s.PostLogin(ctx)

			assert.Equal(t, test.expectedStatus, ctx.Response().Status)
		})
	}
}
