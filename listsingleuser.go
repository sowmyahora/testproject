package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	filter := bson.M{"User_id": userID}

	var User user
	err = collection.FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		log.Println("Error retrieving user:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		log.Println("Error marshaling user to JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
