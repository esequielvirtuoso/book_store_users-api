// Package users service holds the entire business logic to support the users domain.
package users

import (
	"github.com/esequielvirtuoso/book_store_users-api/domain/users"
	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
)

// CreateUser is responsible for getting the input user and writing to the database.
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser is responsible for retriving the user from the database.
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
