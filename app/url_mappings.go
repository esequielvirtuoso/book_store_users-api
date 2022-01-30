package app

import (
	"github.com/esequielvirtuoso/book_store_users-api/controllers/ping"
	"github.com/esequielvirtuoso/book_store_users-api/controllers/users"
)

func mapURLs() {
	// Ping
	// This route allows us to test if the service is up.
	// curl -X GET localhost:8081/ping
	router.GET("/ping", ping.Ping)

	// Get User
	// curl -X GET localhost:8080/users/123 -v
	router.GET("/users/:user_id", users.GetUser)

	// Create User
	// curl -X POST localhost:8081/users
	router.POST("/users", users.CreateUser)
}
