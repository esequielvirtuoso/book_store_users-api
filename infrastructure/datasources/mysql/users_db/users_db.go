// Package usersDb holds the required code to handle database connections.
package usersDb

import (
	"database/sql"
	"log"

	env "github.com/esequielvirtuoso/go_utils_lib/envs"
	"github.com/esequielvirtuoso/go_utils_lib/logger"

	// mysql import
	"github.com/go-sql-driver/mysql"
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
	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
