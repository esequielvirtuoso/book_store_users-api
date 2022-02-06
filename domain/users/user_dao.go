// Package users Data Access Object (DAO) holds the logic to persist or retrieve user to/from a given database.
// DAO is the Access Layer to the database.
package users

import (
	"fmt"

	usersDb "github.com/esequielvirtuoso/book_store_users-api/infrastructure/datasources/mysql/users_db"
	dateUtils "github.com/esequielvirtuoso/book_store_users-api/utils/date_utils"
	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
	mysqlUtils "github.com/esequielvirtuoso/book_store_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get is responsible to retrieve an user from database finding by its id.
func (user *User) Get() *errors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			fmt.Println("error while getting user")
		}
	}()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlUtils.ParseError(getErr)
	}

	return nil
}

// Save is responsible to create a user record within the database.
func (user *User) Save() *errors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			fmt.Println("error while saving user")
		}
	}()

	user.DateCreated = dateUtils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlUtils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlUtils.ParseError(err)
	}
	user.ID = userID
	return nil
}
