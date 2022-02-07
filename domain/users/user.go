// Package users domain holds the definition of users objects and its actions (functions).
// It is a Data Transfer Object (DTO) - object that we will transfer between the persistence layer to the application and backward.
package users

import (
	"strings"

	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

// User struct define the characteristics os a user
type User struct {
	ID           int64  `json:"id"`
	InternalCode int64  `json:"internal_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DateCreated  string `json:"date_created"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	Password     string `json:"password"`
}

// Users is a slice of User
type Users []User

// TrimSpaces is responsible to remove blank spaces from the beginning and the end of user fields values
func (user *User) TrimSpaces() {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)
}

// Validate is responsible to verify if the values assigned to the user's characteristics are allowed
func (user *User) Validate() *errors.RestErr {
	user.TrimSpaces()

	if user.InternalCode <= 0 {
		return errors.NewBadRequestError("invalid internal code")
	}

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email addess")
	}

	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid first name")
	}

	if user.LastName == "" {
		return errors.NewBadRequestError("invalid last name")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
