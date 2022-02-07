// Package users domain holds the definition of users objects and its actions (functions).
// It is a Data Transfer Object (DTO) - object that we will transfer between the persistence layer to the application and backward.
package users

import (
	"strings"

	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
	"github.com/esequielvirtuoso/book_store_users-api/utils/regex"
)

const (
	// StatusActive defines the default user status
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

// validatePassword validates if password is compliant to the minimal security rules.
// It returns False and a message when password is not valid.
func validatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "password must have at least 8 characters"
	}

	patterns := []string{"[A-Z]:password must have at least one upper case character:t",
		"[a-z]:password must have at least one lower case character:t",
		"[0-9]:password must have at least one number:t",
		"[^A-Za-z0-9]:password must have at least one special character:t",
		"[\t\n\f\r ]:password cannot have white spaces:f"}

	validations := regex.DefineRegexPatternsAndMessages(patterns)

	for _, validation := range validations {
		if validation.Found {
			found, _ := regex.AssertRegexPattern(password, validation.Regex)
			if !found {
				return false, validation.Message
			}
		} else {
			found, _ := regex.AssertRegexPattern(password, validation.Regex)
			if found {
				return false, validation.Message
			}
		}
	}
	return true, ""
}

// TrimSpaces is responsible to remove blank spaces from the beginning and the end of user fields values
func (user *User) TrimSpaces() {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
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

	passIsValid, message := validatePassword(user.Password)
	if !passIsValid {
		return errors.NewBadRequestError(message)
	}

	return nil
}
