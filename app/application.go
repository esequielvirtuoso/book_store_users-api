// Package app starts the application and map the HTTP routes.
package app

import (
	"github.com/gin-gonic/gin"
)

var (
	// NOTE: This is the only layer we are defining and using the HTTP server
	router = gin.Default()
)

// StartApplication attempts to map the API routes
func StartApplication() {
	mapURLs()
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}
