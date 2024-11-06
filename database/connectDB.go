package database

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func InitDB() *sql.DB {
	var err error
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return Db
}
