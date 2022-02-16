// Package users Data Access Object (DAO) holds the logic to persist or retrieve user to/from a given database.
// DAO is the Access Layer to the database.
package users

import (
	"fmt"
	"strings"

	usersDb "github.com/esequielvirtuoso/book_store_users-api/infrastructure/datasources/mysql/users_db"
	dateUtils "github.com/esequielvirtuoso/go_utils_lib/date"
	"github.com/esequielvirtuoso/go_utils_lib/logger"
	mysqlUtils "github.com/esequielvirtuoso/go_utils_lib/mysql_errors"
	restErrors "github.com/esequielvirtuoso/go_utils_lib/rest_errors"
)

const (
	queryInsertUser             = "INSERT INTO users(internal_code, first_name, last_name, email, status, password, date_created, updated_at) VALUES(?,?,?,?,?,?,?,?);"
	queryGetUser                = "SELECT id, internal_code, first_name, last_name, email, status, date_created, updated_at FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET internal_code=?, first_name=?, last_name=?, email=?, status=?, password=?, updated_at=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus       = "SELECT id, internal_code, first_name, last_name, email, status, date_created, updated_at FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, internal_code, first_name, last_name, email, status, date_created, updated_at FROM users WHERE email=? AND password=? AND status=?;"
)

// Get is responsible to retrieve an user from database finding by its id.
func (user *User) Get() restErrors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return restErrors.NewInternalServerError("database error")
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error("error while defering getting user statement", err)
		}
	}()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.InternalCode, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated, &user.UpdatedAt); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return restErrors.NewInternalServerError("database error")
	}

	return nil
}

// Save is responsible to create a user record within the database.
func (user *User) Save() restErrors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return restErrors.NewInternalServerError("database error")
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error("error while defering saving user statement", err)
		}
	}()

	insertResult, saveErr := stmt.Exec(user.InternalCode, user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated, user.UpdatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return restErrors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return restErrors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}

// Update is responsible to update a user record within the database.
func (user *User) Update() restErrors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return restErrors.NewInternalServerError("database error")
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error("error while defering updating user statement", err)
		}
	}()

	_, err = stmt.Exec(user.InternalCode, user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		return mysqlUtils.ParseError(err)
	}
	return nil
}

// Delete is responsible to delete a user record from the database.
func (user *User) Delete() restErrors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return restErrors.NewInternalServerError("database error")
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error("error while defering delete user statement", err)
		}
	}()

	user.UpdatedAt = dateUtils.GetNowString()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		return restErrors.NewInternalServerError("database error")
	}
	return nil
}

// FindByStatus is responsible to find users by status.
func (user *User) FindByStatus(status string) ([]User, restErrors.RestErr) {
	stmt, err := usersDb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, restErrors.NewInternalServerError("database error")
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error("error while defering find users by status statement", err)
		}
	}()

	rows, errFindByStatus := stmt.Query(status)
	// avoid keeping open connections to the database
	defer func() {
		err = rows.Close()
		if err != nil {
			logger.Error("error while closinf find users by status result rows", err)
		}
	}()

	if errFindByStatus != nil {
		logger.Error("error when trying to find user by status", errFindByStatus)
		return nil, restErrors.NewInternalServerError("database error")
	}

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.InternalCode, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated, &user.UpdatedAt); err != nil {
			logger.Error("error when trying to scan users returned rows", err)
			return nil, restErrors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		logger.Info(fmt.Sprintf("no users matching status %s", status))
		return nil, restErrors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil

}

// FindByEmailAndPassword is responsible to retrieve an user from database finding by its email and password.
func (user *User) FindByEmailAndPassword() restErrors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return restErrors.NewInternalServerError("database error")
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error("error while defering getting user by email and password statement", err)
		}
	}()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.InternalCode, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated, &user.UpdatedAt); getErr != nil {
		if strings.Contains(getErr.Error(), mysqlUtils.ErrorNoRows) {
			return restErrors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return restErrors.NewInternalServerError("database error")
	}

	return nil
}
