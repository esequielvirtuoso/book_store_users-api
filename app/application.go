// Package app starts the application and map the HTTP routes.
package app

import (
	env "github.com/esequielvirtuoso/go_utils_lib/envs"
	"github.com/esequielvirtuoso/go_utils_lib/logger"
	"github.com/gin-gonic/gin"
)

const (
	mysqlURL        = "MYSQL_URL"
	defaultmysqlURL = "root:passwd@tcp(127.0.0.1:3305)/users_db?charset=utf8"
)

var (
	// NOTE: This is the only layer we are defining and using the HTTP server
	router = gin.Default()
)

// StartApplication attempts to map the API routes
func StartApplication() {
	env.CheckRequired(mysqlURL)

	mapURLs()

	logger.Info("about to start the users application")
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}

func getMySQLURL() string {
	return env.GetString(mysqlURL, defaultmysqlURL)
}
