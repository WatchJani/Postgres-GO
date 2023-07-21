package db

import (
	"database/sql"
	"log"
)

var Conn *sql.DB

func Connect() {
	var err error
	Conn, err = sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
}
