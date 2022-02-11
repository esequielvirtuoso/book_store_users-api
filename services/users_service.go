// Package services service holds the entire business logic to support the users domain.
package services

import (
	"github.com/esequielvirtuoso/book_store_users-api/domain/users"
	cryptoUtils "github.com/esequielvirtuoso/book_store_users-api/utils/crypto_utils"
	dateUtils "github.com/esequielvirtuoso/book_store_users-api/utils/date_utils"
	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
)

var (
	// UsersService is a public variable to export the users service interface
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(*users.LoginRequest) (*users.User, *errors.RestErr)
}

// CreateUser is responsible for getting the input user and writing to the database.
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateUtils.GetNowDbFormat()
	user.UpdatedAt = dateUtils.GetNowDbFormat()
	user.Password = cryptoUtils.GetSha256(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser is responsible for retriving the user from the database.
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser is responsible for updating the users records
func (s *usersService) UpdateUser(isPartialUpdate bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartialUpdate {
		user.TrimSpaces()
		if user.InternalCode >= 0 {
			currentUser.InternalCode = user.InternalCode
		}

		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}

		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}

		if user.Email != "" {
			currentUser.Email = user.Email
		}

		if user.Status != "" {
			currentUser.Status = user.Status
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}
		currentUser.InternalCode = user.InternalCode
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
		currentUser.Status = user.Status
		currentUser.Password = user.Password
	}

	currentUser.UpdatedAt = dateUtils.GetNowDbFormat()
	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

// DeleteUser is responsible for deleting the user
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

// Search is responsible for finding users by its characteristics
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request *users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: cryptoUtils.GetSha256(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
