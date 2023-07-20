package main

import (
	"context"
	"encoding/json"
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

type user struct {
	User_id int64    `bson:"User_id"`
	Name    string   `bson:"Name"`
	Phone   string   `bson:"Phone"`
	Address address  `bson:"Address"`
	Hobbies []string `bson:"Hobbies"`
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	var User user

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if User.Name == "" {
		http.Error(w, "Invalid user: Name cannot be empty", http.StatusBadRequest)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	_, err = collection.InsertOne(context.TODO(), User)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
