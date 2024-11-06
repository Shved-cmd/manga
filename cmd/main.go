package main

import (
	"go.mod/database"
)
import _ "github.com/lib/pq"

var DB = database.InitDB()

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	query := `insert into users (username, password) values ($1,$2)`
	DB.Exec(query, "new_user", "secure password")
}
