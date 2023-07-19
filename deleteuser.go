package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userID := params.Get("id")
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	filter := bson.M{"User_id": userIDInt}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if result.DeletedCount == 0 {
		http.NotFound(w, r)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}
