package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetUserPassword(t *testing.T) {
	tests := []struct {
		name        string
		input       GetUserPasswordInput
		beforeTest  func(sqlmock.Sqlmock)
		expectedErr error
		output      GetUserPasswordOutput
	}{
		{
			name: "Successful Get",
			input: GetUserPasswordInput{
				PhoneNumber: "1234566",
			},
			output: GetUserPasswordOutput{
				UserID:       1,
				Password:     "hashed",
				PasswordSalt: "salt",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery("SELECT user_id, password_hash, password_salt FROM users WHERE phone_number = $1").
					WithArgs("1234566").
					WillReturnRows(sqlmock.NewRows(
						[]string{"user_id", "password_hash", "password_salt"}).
						AddRow(1, "hashed", "salt"),
					)
			},
			expectedErr: nil,
		},
		{
			name: "error",
			input: GetUserPasswordInput{
				PhoneNumber: "1234566",
			},
			output: GetUserPasswordOutput{},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery("SELECT user_id, password_hash, password_salt FROM users WHERE phone_number = $1").
					WithArgs("1234566").
					WillReturnError(errors.New("error sample"))
			},
			expectedErr: errors.New("error sample"),
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()

			r := Repository{
				Db: db,
			}
			if test.beforeTest != nil {
				test.beforeTest(mock)
			}
			output, err := r.GetUserPassword(context.Background(), test.input)

			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.output, output)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
func TestRepository_GetUserByUserID(t *testing.T) {
	tests := []struct {
		name        string
		input       GetUserByUserIDInput
		beforeTest  func(sqlmock.Sqlmock)
		expectedErr error
		output      GetUserByUserIDOutput
	}{
		{
			name: "Successful Get",
			input: GetUserByUserIDInput{
				UserID: 1,
			},
			output: GetUserByUserIDOutput{
				PhoneNumber: "123456789",
				Fullname:    "Updated Name",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery("SELECT phone_number,full_name FROM users WHERE user_id = $1").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows(
						[]string{"phone_number", "full_name"}).
						AddRow("123456789", "Updated Name"),
					)
			},
			expectedErr: nil,
		},
		{
			name: "error",
			input: GetUserByUserIDInput{
				UserID: 2,
			},
			output: GetUserByUserIDOutput{},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery("SELECT phone_number,full_name FROM users WHERE user_id = $1").
					WithArgs(2).
					WillReturnError(errors.New("error sample"))
			},
			expectedErr: errors.New("error sample"),
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()

			r := Repository{
				Db: db,
			}
			if test.beforeTest != nil {
				test.beforeTest(mock)
			}
			output, err := r.GetUserByUserID(context.Background(), test.input)

			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.output, output)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_UpdateUserByUserID(t *testing.T) {
	tests := []struct {
		name        string
		input       UpdateUserByUserIDInput
		beforeTest  func(sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name: "Successful Update",
			input: UpdateUserByUserIDInput{
				UserID:      1,
				PhoneNumber: "123456789",
				Fullname:    "Updated Name",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec("UPDATE users SET phone_number = $1, full_name = $2 WHERE user_id = $3").
					WithArgs("123456789", "Updated Name", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: nil,
		},
		{
			name: "Phone Number Conflict",
			input: UpdateUserByUserIDInput{
				UserID:      2,
				PhoneNumber: "987654321",
				Fullname:    "Updated Name",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec("UPDATE users SET phone_number = $1, full_name = $2 WHERE user_id = $3").
					WithArgs("987654321", "Updated Name", 2).
					WillReturnError(&pq.Error{Code: "23505", Message: "duplicate key value"})
			},
			expectedErr: &pq.Error{Code: "23505", Message: "duplicate key value"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()

			r := Repository{
				Db: db,
			}
			if test.beforeTest != nil {
				test.beforeTest(mock)
			}
			err = r.UpdateUserByUserID(context.Background(), test.input)

			assert.Equal(t, test.expectedErr, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_UpdateSuccessLoginByUserID(t *testing.T) {
	tests := []struct {
		name        string
		input       UpdateSuccessLoginByUserIDInput
		beforeTest  func(sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name: "Successful Update",
			input: UpdateSuccessLoginByUserIDInput{
				PhoneNumber: "123456789",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec("UPDATE users SET login_attempts = login_attempts + 1 WHERE phone_number = $1").
					WithArgs("123456789").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: nil,
		},
		{
			name: "Phone Number Error",
			input: UpdateSuccessLoginByUserIDInput{
				PhoneNumber: "987654321",
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectExec("UPDATE users SET login_attempts = login_attempts + 1 WHERE phone_number = $1").
					WithArgs("987654321").
					WillReturnError(errors.New("some error"))
			},
			expectedErr: errors.New("some error"),
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()

			r := Repository{
				Db: db,
			}
			if test.beforeTest != nil {
				test.beforeTest(mock)
			}
			err = r.UpdateSuccessLoginByUserID(context.Background(), test.input)

			assert.Equal(t, test.expectedErr, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
