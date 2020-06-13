package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_user_username = "mysql_user_username"
	mysql_user_password = "mysql_user_password"
	mysql_user_host     = "mysql_user_host"
	mysql_user_schema   = "mysql_user_schema"
)

var (
	Client *sql.DB

	userName = os.Getenv("mysql_user_username")
	password = os.Getenv("mysql_user_password")
	host     = os.Getenv("mysql_user_host")
	schema   = os.Getenv("mysql_user_schema")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		userName, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connection made succesfully")
	fmt.Println("Connection made succesfully")
}
