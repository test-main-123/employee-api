package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	store := NewEmployeeStore()

	r := mux.NewRouter()
	r.HandleFunc("/employees", store.ListEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", store.GetEmployee).Methods("GET")
	r.HandleFunc("/employees", store.CreateEmployee).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
