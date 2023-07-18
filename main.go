package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func insertUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	profilepic, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read profile image: "+err.Error(), http.StatusBadRequest)
		return
	}

	user.ProfilePicture = profilepic

	collection := client.Database("testdb").Collection("users")

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func main() {

	http.HandleFunc("/users/insert", insertUser)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
