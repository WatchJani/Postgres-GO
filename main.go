package main

import (
	"net/http"
	"root/db"
	"root/routes"

	_ "github.com/lib/pq"
)

func main() {
	db.Connect()

	http.HandleFunc("/users", routes.CreateUser)
	http.HandleFunc("/string", routes.String)
	http.ListenAndServe(":8080", nil)
}
