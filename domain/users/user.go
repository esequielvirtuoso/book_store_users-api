// Package users domain holds the definition of users objects and its actions (functions).
package users

import (
	"strings"

	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
)

// User struct define the characteristics os a user
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate is responsible to verify if the values assigned to the user's characteristics are allowed
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email addess")
	}
	return nil
}
