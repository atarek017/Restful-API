package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var db *sql.DB
var err error

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitDatabase() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users [] User
	result, err := db.Query("select * from users")
	if err != nil {
		println(err.Error())
		panic(err.Error())
	}
	for result.Next() {
		var user User
		err = result.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
	defer result.Close()
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = rand.Intn(1000000) //Mock ID -not safe
	insert, err := db.Query("INSERT INTO users VALUES ('" + strconv.Itoa(user.ID) + "', '" + user.Name + "', '" + user.Email + "')")
	if err != nil {
		panic(err.Error())
		println(err.Error())
	}
	defer insert.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get  params

	delete, err := db.Query("delete  FROM shop.users where id='" + params["id"] + "'")
	if err != nil {
		print(err.Error())
		print(err.Error())
	}
	defer delete.Close()
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	insert, err := db.Query("UPDATE `shop`.`users` SET `name`='" + user.Name + " ', `email`='" + user.Email + "' WHERE `id`=  ' " + strconv.Itoa(user.ID) + "'")
	if err != nil {
		panic(err.Error())
		println(err.Error())
	}
	defer insert.Close()
}
