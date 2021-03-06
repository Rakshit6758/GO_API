package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() { // this method will initialize the router
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r)) // port for the server ios 9000

}

func main() {
	InitialMigration()
	initializeRouter()

}
