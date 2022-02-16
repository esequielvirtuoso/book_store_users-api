// Package users marshall holds the logic to build up the returns users objects to avoid returning sensitive information to public HTTP requests.
package users

import (
	"encoding/json"

	"github.com/esequielvirtuoso/go_utils_lib/logger"
)

// PublicUser struct define the characteristics of a user that we can return to public HTTP requests
type PublicUser struct {
	ID           int64  `json:"id"`
	InternalCode int64  `json:"internal_code"`
	DateCreated  string `json:"date_created"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
}

// PrivateUser struct define the characteristics of a user that we can return to private (authenticated) HTTP requests
type PrivateUser struct {
	ID           int64  `json:"id"`
	InternalCode int64  `json:"internal_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DateCreated  string `json:"date_created"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
}

// Marshall is responsible to marshall the result users when it is a search and returns a slice
func (users Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}

	return result
}

// Marshall is responsible to marshall the result user
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:           user.ID,
			InternalCode: user.InternalCode,
			DateCreated:  user.DateCreated,
			UpdatedAt:    user.UpdatedAt,
			Status:       user.Status,
		}
	}
	userJSON, errMarshal := json.Marshal(user)
	if errMarshal != nil {
		logger.Error("error while marshalling user", errMarshal)
	}
	var privateUser PrivateUser
	if errUnmarshal := json.Unmarshal(userJSON, &privateUser); errUnmarshal != nil {
		logger.Error("error while unmarshalling user", errUnmarshal)
	}

	return privateUser
}
