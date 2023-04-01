package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"root/db"
)

type Dev struct {
	ID     uint
	Name   string
	Age    uint16
	City   string
	Salary float64
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query("SELECT * FROM staff.dev")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var devs []Dev
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
