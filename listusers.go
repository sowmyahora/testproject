package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func listUsers(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error retrieving users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var Users []user
	for cursor.Next(context.TODO()) {
		var User user
		err := cursor.Decode(&User)
		if err != nil {
			log.Println("Error decoding user:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		Users = append(Users, User)
	}

	usersJSON, err := json.Marshal(Users)
	if err != nil {
		log.Println("Error marshaling users to JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}
