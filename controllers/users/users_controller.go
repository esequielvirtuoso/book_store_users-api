// Package users provide the functionalities or the entry points to interact with the users API.
// Take the request, validate if we have all the parameters that we need to
// handle this request and send this handling to the service where we have
// the required business logic.
package users

import (
	"net/http"
	"strconv"

	"github.com/esequielvirtuoso/book_store_users-api/domain/users"
	users_service "github.com/esequielvirtuoso/book_store_users-api/services/users"
	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of users requests.
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid body request")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := users_service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser handles the get users requests.
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := users_service.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
