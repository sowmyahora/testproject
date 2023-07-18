package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users/insert", insertUser)
	http.HandleFunc("/users/update", updateUser)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
