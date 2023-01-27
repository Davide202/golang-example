package users_db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang-users/utils/logger"
	"log"
	_ "os"
)

/*
const (

	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"

)
*/
var (
	Client *sql.DB
	/*
		username = os.Getenv(mysqlUsersUsername)
		password = os.Getenv(mysqlUsersPassword)
		host     = os.Getenv(mysqlUsersHost)
		schema   = os.Getenv(mysqlUsersSchema)
	*/
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "admin", "localhost", "db_users",
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
