package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//http.HandleFunc("/users/insert", handleUser)
	//http.HandleFunc("/users/insert", handleUser)
	r := mux.NewRouter()
	r.HandleFunc("/users", insertUser).Methods("POST")
	r.HandleFunc("/users/{user_id}", updateUser).Methods("PUT")
	r.HandleFunc("/users", listUsers).Methods("GET")
	r.HandleFunc("/users/{user_id}", getUser).Methods("GET")
	r.HandleFunc("/users/{user_id}", deleteUser).Methods("DELETE")
	http.Handle("/", r)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
