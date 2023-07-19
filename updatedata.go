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

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")
	filter := bson.M{"User_id": int64(user.User_id)}

	var existingUser User
	err = collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		log.Println("Error retrieving user:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Printf("Existing user: %+v", existingUser)

	update := bson.M{"$set": bson.M{
		"Name":    user.Name,
		"Phone":   user.Phone,
		"Address": user.Address,
		"Hobbies": user.Hobbies,
	}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}
