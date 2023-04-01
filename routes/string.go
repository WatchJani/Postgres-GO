package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"root/db"
)

func String(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()
		fmt.Println("Received message:", string(body))

		rows, err := db.Conn.Query("SELECT name FROM staff.dev WHERE name LIKE $1", string(body)+"%")

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var devs []string
		for rows.Next() {
			var dev string
			if err := rows.Scan(&dev); err != nil {
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
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
