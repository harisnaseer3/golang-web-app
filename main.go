package main

import (
	"net/http"
	"golang-web-app/handlers"
	"golang-web-app/db"
)

func main() {
	db.Init()

	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/patients", handlers.GetPatients)

	http.ListenAndServe(":8080", nil)
}
