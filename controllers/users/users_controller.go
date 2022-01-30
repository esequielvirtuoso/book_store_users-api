// Provide the functionalities or the entry points to interact with the users API.
// Take the request, validate if we have all the parameters that we need to
// handle this request and send this handling to the service where we have
// the required business logic.
package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
