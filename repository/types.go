// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type UserInput struct {
	PhoneNumber  string
	Fullname     string
	Password     string
	PasswordSalt string
}

type UserOutput struct {
	ID int64
}

type GetUserPasswordInput struct {
	PhoneNumber string
}

type GetUserPasswordOutput struct {
	UserID       int64
	Password     string
	PasswordSalt string
}

type GetUserByUserIDInput struct {
	UserID int64
}

type GetUserByUserIDOutput struct {
	PhoneNumber string
	Fullname    string
}

type UpdateUserByUserIDInput struct {
	UserID      int64
	Fullname    string
	PhoneNumber string
}

type UpdateSuccessLoginByUserIDInput struct {
	PhoneNumber string
}
