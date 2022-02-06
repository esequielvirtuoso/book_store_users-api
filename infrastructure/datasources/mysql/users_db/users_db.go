// Package usersDb holds the required code to handle database connections.
package usersDb

import (
	"database/sql"
	"log"

	"github.com/esequielvirtuoso/book_store_users-api/utils/env"
	// mysql import
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlURL = "MYSQL_URL" // root:passwd@tcp(127.0.0.1:3305)/users_db?charset=utf8
)

var (
	// Client variable is a client connection to the MySQL database.
	Client *sql.DB
)

// init is the entrypoint of this package. It stablish a client connection to the MySQL database.
func init() {
	env.CheckRequired(mysqlURL)

	var err error
	Client, err = sql.Open("mysql", env.GetString(mysqlURL))
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
