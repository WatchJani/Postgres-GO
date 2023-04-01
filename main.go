package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Dev struct {
	ID     uint
	Name   string
	Age    uint16
	City   string
	Salary float64
}

func init() {
	var err error
	db, err = sql.Open("postgres", "user=janko dbname=company password=JankoKondic72621@ sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/users", createUser)
	http.ListenAndServe(":8080", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM staff.dev")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	devs := make([]Dev, 0)
	for rows.Next() {
		var dev Dev
		if err := rows.Scan(&dev.ID, &dev.Name, &dev.Age, &dev.City, &dev.Salary); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		devs = append(devs, dev)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kodiranje slice-a korisnika u JSON i slanje kao HTTP odgovor
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(devs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(devs)
}
