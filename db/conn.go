package db

import (
	"database/sql"
	"log"
)

var Conn *sql.DB

func Connect() {
	var err error
	Conn, err = sql.Open("postgres", "user=janko dbname=company password=JankoKondic72621@ sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
