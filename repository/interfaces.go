// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	InsertUser(ctx context.Context, input UserInput) (output UserOutput, err error)
	GetUserPassword(ctx context.Context, input GetUserPasswordInput) (output GetUserPasswordOutput, err error)
	UpdateSuccessLoginByUserID(ctx context.Context, input UpdateSuccessLoginByUserIDInput) (err error)
	GetUserByUserID(ctx context.Context, input GetUserByUserIDInput) (output GetUserByUserIDOutput, err error)
	UpdateUserByUserID(ctx context.Context, input UpdateUserByUserIDInput) (err error)
}
