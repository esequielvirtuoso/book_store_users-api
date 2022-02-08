// Package app map the HTTP routes.
package app

import (
	"github.com/esequielvirtuoso/book_store_users-api/controllers/ping"
	"github.com/esequielvirtuoso/book_store_users-api/controllers/users"
)

// mapURLs map the HTTP routes.
func mapURLs() {
	// Ping
	// This route allows us to test if the service is up.
	// curl -X GET localhost:5001/ping
	router.GET("/ping", ping.Ping)

	// Get User
	// curl -X GET localhost:5001/users/123 -v
	router.GET("/users/:user_id", users.Get)

	// Create User
	// curl -X POST localhost:5001/users
	router.POST("/users", users.Create)

	// Full Update User
	// curl -X PUT localhost:5001/users/123
	router.PUT("/users/:user_id", users.Update)

	// Partial Update User
	// curl -X PATCH localhost:5001/users/123
	router.PATCH("/users/:user_id", users.Update)

	// Delete User
	// curl -X DELETE localhost:5001/users/123
	router.DELETE("/users/:user_id", users.Delete)

	// Search Users
	// curl -X GET localhost:5001/internal/users/search
	router.GET("/internal/users/search", users.Search)
}
