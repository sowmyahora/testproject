package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users/insert", insertUser)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
