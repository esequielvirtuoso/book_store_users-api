package app

import "github.com/gin-gonic/gin"

var (
	// NOTE: This is the only layer we are defining and using the HTTP server
	router = gin.Default()
)

func StartApplication() {
	mapURLs()
	router.Run(":8081")
}
