package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func (r *Repository) InsertUser(ctx context.Context, input UserInput) (output UserOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO users (phone_number, full_name, password_hash, password_salt) VALUES ($1, $2, $3, $4) returning user_id",
		input.PhoneNumber,
		input.Fullname,
		input.Password,
		input.PasswordSalt).Scan(&output.ID)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserPassword(ctx context.Context, input GetUserPasswordInput) (output GetUserPasswordOutput, err error) {
	query := "SELECT user_id, password_hash, password_salt FROM users WHERE phone_number = $1"
	row := r.Db.QueryRowContext(ctx, query, input.PhoneNumber)

	err = row.Scan(&output.UserID, &output.Password, &output.PasswordSalt)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (r *Repository) UpdateSuccessLoginByUserID(ctx context.Context, input UpdateSuccessLoginByUserIDInput) (err error) {
	query := "UPDATE users SET login_attempts = login_attempts + 1 WHERE phone_number = $1"
	_, err = r.Db.ExecContext(ctx, query, input.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByUserID(ctx context.Context, input GetUserByUserIDInput) (output GetUserByUserIDOutput, err error) {
	query := "SELECT phone_number,full_name FROM users WHERE user_id = $1"
	row := r.Db.QueryRowContext(ctx, query, input.UserID)
	err = row.Scan(&output.PhoneNumber, &output.Fullname)
	if err != nil {
		return output, err
	}

	return output, nil
}

func (r *Repository) UpdateUserByUserID(ctx context.Context, input UpdateUserByUserIDInput) (err error) {
	// Construct the SQL query and arguments
	query := "UPDATE users SET"
	args := []interface{}{}
	if input.PhoneNumber != "" {
		query += " phone_number = $1,"
		args = append(args, input.PhoneNumber)
	}
	if input.Fullname != "" {
		query += " full_name = $2"
		args = append(args, input.Fullname)
	} else {
		// Remove the trailing comma from the query
		query = strings.TrimSuffix(query, ",")
	}
	query += " WHERE user_id = $3"
	args = append(args, input.UserID)

	_, err = r.Db.Exec(query, args...)

	// Check for unique constraint violation error
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" && strings.Contains(pqErr.Message, "phone_number") {
			return fmt.Errorf("phone number already exists")
		}
	}
	return
}
