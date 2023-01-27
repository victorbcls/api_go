package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type User struct {
	UserId   int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
	BirthDay string `json:"birthday"`
}

func GetUsers(Client *sql.DB) ([]byte, error) {
	var Users = []User{}

	rows, err := Client.Query("SELECT Username, Fullname, Birthday FROM Go.Users;")
	for rows.Next() {
		var usr User
		err := rows.Scan(&usr.Username, &usr.FullName, &usr.BirthDay)
		if err != nil {
			fmt.Println(err)
		}
		Users = append(Users, usr)
	}
	fmt.Println(Users)
	response, err := json.Marshal(Users)

	return response, err

}
func GetUserByUsername(Client *sql.DB, username string) ([]byte, error) {
	var User User
	rows := Client.QueryRow("SELECT Username, Fullname, Birthday from Go.Users where Username = ?", username)
	fmt.Println(rows)
	err := rows.Scan(&User.Username, &User.FullName, &User.BirthDay)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	response, err := json.Marshal(User)
	return response, err

}
