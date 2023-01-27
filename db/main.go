package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

func Connect() (*sql.DB, error) {
	connection, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/Go")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")
	Client = connection
	return Client, err

}

func Query(query string) sql.Result {
	response, err := Client.Exec(query)

	if err != nil {
		fmt.Println(err)
	}
	return response

}
