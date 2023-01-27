package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	helloworld "github.com/victorbcls/api-go/controllers/hello-world"
	"github.com/victorbcls/api-go/controllers/users"

	"github.com/victorbcls/api-go/db"
)

func api() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloworld.HelloWord)
	r.HandleFunc("/users", getAllUsers)
	r.HandleFunc("/user/{username}", getUserByUsername).Methods("GET")

	http.ListenAndServe(":8000", r)
}

func getAllUsers(res http.ResponseWriter, req *http.Request) {
	response, err := users.GetUsers(Client)
	fmt.Println(string(response))
	if err != nil {
		response, _ := json.Marshal(helloworld.DefaultResponse{Status: 500, Message: "Unknown Error"})
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		res.Write(response)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(response)
	}

}

func getUserByUsername(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]
	response, err := users.GetUserByUsername(Client, username)
	if err != nil {
		response, _ := json.Marshal(helloworld.DefaultResponse{Status: 500, Message: "Unknown Error"})
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		res.Write(response)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(response)
	}
}

var Client *sql.DB

func main() {
	fmt.Println("Rest Api")
	client, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}
	Client = client
	api()
}
