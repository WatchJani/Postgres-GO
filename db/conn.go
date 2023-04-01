package db

import (
	"database/sql"
	"fmt"
)

func Connect() (*sql.DB, error) {
	connStr := "user=janko dbname=company password=JankoKondic72621@ sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %v", err)
	}
	
	return db, nil
}
