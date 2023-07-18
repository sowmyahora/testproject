package main

import (
	"log"
	"net/http"
)

type address struct {
	Street  string `bson:"Street"`
	City    string `bson:"City"`
	State   string `bson:"State"`
	Country string `bson:"Country"`
}

type User struct {
	User_id        int64    `bson:"User_id"`
	Name           string   `bson:"Name"`
	Phone          string   `bson:"Phone"`
	Address        address  `bson:"Address"`
	Hobbies        []string `bson:"Hobbies"`
	ProfilePicture []byte   `bson:"ProfilePicture"`
}

func main() {

	http.HandleFunc("/users/insert", insertUser)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
