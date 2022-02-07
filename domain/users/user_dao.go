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
	queryInsertUser       = "INSERT INTO users(internal_code, first_name, last_name, email, status, password, date_created, updated_at) VALUES(?,?,?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, internal_code, first_name, last_name, email, status, date_created, updated_at FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET internal_code=?, first_name=?, last_name=?, email=?, status=?, password=?, updated_at=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, internal_code, first_name, last_name, email, status, date_created, updated_at FROM users WHERE status=?;"
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
	if getErr := result.Scan(&user.ID, &user.InternalCode, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated, &user.UpdatedAt); getErr != nil {
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

	insertResult, saveErr := stmt.Exec(user.InternalCode, user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated, user.UpdatedAt)
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

// Update is responsible to update a user record within the database.
func (user *User) Update() *errors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			fmt.Println("error while updating user")
		}
	}()

	_, err = stmt.Exec(user.InternalCode, user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		return mysqlUtils.ParseError(err)
	}
	return nil
}

// Delete is responsible to delete a user record from the database.
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersDb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			fmt.Println("error while deleting user")
		}
	}()

	user.UpdatedAt = dateUtils.GetNowString()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysqlUtils.ParseError(err)
	}
	return nil
}

// FindByStatus is responsible to find users by status.
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersDb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			fmt.Println("error while finding users by status")
		}
	}()

	rows, errFindById := stmt.Query(status)
	// avoid keeping open connections to the database
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("error while finding users by status")
		}
	}()

	if errFindById != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.InternalCode, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated, &user.UpdatedAt); err != nil {
			return nil, mysqlUtils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil

}
