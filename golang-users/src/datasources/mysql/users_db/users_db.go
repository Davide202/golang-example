package users_db

import (
	"database/sql"
	"fmt"
	"golang-users/utils/logger"
	"log"
	"os"
	_ "os"

	"github.com/go-sql-driver/mysql"
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
	var dataSourceName string
	host, maybe := os.LookupEnv("mysql_users_host")
	if maybe {
		log.Println("Evn " + "mysql_users_host=" + host + " is present")
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "admin", host, "db_users")
	} else {
		log.Println("Evn " + "mysql_users_host" + " is NOT present")
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "admin", "localhost", "db_users")
	}

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("database NOT successfully configured")
	}
	if err = Client.Ping(); err != nil {
		log.Println("database NOT successfully configured")
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
