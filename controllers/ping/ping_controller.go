// Package ping allows the cloud services and clients to verify if the application is running.
package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Ping function is responsible to notify the clients or cloud services
*  if the application is up and running
 */
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Users application up and running")
}
