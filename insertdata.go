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
	street  string `bson:"Street"`
	city    string `bson:"City"`
	state   string `bson:"State"`
	country string `bson:"Country"`
}

type User struct {
	user_id int64    `bson:"User_id"`
	name    string   `bson:"Name"`
	phone   string   `bson:"Phone"`
	address address  `bson:"Address"`
	hobbies []string `bson:"Hobbies"`
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if user.name == "" {
		http.Error(w, "Invalid user: Name cannot be empty", http.StatusBadRequest)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
