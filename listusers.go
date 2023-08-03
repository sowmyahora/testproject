package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func listUsers(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		response := httpresponse{
			Message: "Error retrieving users:",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}
	defer cursor.Close(ctx)

	var Users []user
	for cursor.Next(ctx) {
		var User user
		err := cursor.Decode(&User)
		if err != nil {
			response := httpresponse{
				Message: "error",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

			return

		}
		Users = append(Users, User)
	}

	usersJSON, err := json.Marshal(Users)
	if err != nil {
		response := httpresponse{
			Message: "Error marshaling users to JSON:",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}
