package main

import (
	"log"
	"net/http"

	"testGo/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/alert_policy", handlers.CreateAlertPolicy).Methods("POST")
	r.HandleFunc("/alert_policy/{id}", handlers.UpdateAlertPolicy).Methods("PUT")
	r.HandleFunc("/alert_policy/{id}", handlers.DeleteAlertPolicy).Methods("DELETE")
	r.HandleFunc("/alert_policy/{id}", handlers.FetchAlertPolicy).Methods("GET")

	http.Handle("/", r)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
