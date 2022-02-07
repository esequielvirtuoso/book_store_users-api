// Package users provide the functionalities or the entry points to interact with the users API.
// Take the request, validate if we have all the parameters that we need to
// handle this request and send this handling to the service where we have
// the required business logic.
package users

import (
	"net/http"
	"strconv"

	"github.com/esequielvirtuoso/book_store_users-api/domain/users"
	"github.com/esequielvirtuoso/book_store_users-api/services"
	"github.com/esequielvirtuoso/book_store_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

const (
	isPublic = "true"
)

// getUserID isolate the get user id action to be reused
func getUserID(inputUserID string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(inputUserID, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")

	}
	return userID, nil
}

// Create handles the creation of users requests.
func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid body request")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == isPublic))
}

// Get handles the get users requests.
func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == isPublic))
}

// Update handles the update of users requests.
func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid body request")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	isPartialUpdate := c.Request.Method == http.MethodPatch

	result, errUpdate := services.UsersService.UpdateUser(isPartialUpdate, user)

	if errUpdate != nil {
		c.JSON(errUpdate.Status, errUpdate)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == isPublic))
}

// Delete handles the update of users requests.
func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if errDelete := services.UsersService.DeleteUser(userID); errDelete != nil {
		c.JSON(errDelete.Status, errDelete)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

// Search handles find users by its characteristics
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == isPublic))
}
