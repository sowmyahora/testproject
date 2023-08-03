package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func updateUser(w http.ResponseWriter, r *http.Request) {

	userIDstr := mux.Vars(r)["user_id"]
	if userIDstr == "" {
		response := httpresponse{
			Message: "Missing 'user_id' query parameter",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		response := httpresponse{
			Message: "Invalid 'user_id' query parameter",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	fmt.Println(userID, "user id retrieved")
	User, err := validateUpdateRequest(r)
	if err != nil {
		response := httpresponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("testdb").Collection("users")
	filter := bson.M{"user_id": userID}

	userToUpdate := make(map[string]interface{})

	if User.Name != "" {
		userToUpdate["name"] = User.Name
	}

	if User.Phone != "" {
		userToUpdate["phone"] = User.Phone
	}

	if User.Address.Street != "" {
		userToUpdate["address.street"] = User.Address.Street
	}
	if User.Address.City != "" {
		userToUpdate["address.city"] = User.Address.City
	}
	if User.Address.State != "" {
		userToUpdate["address.state"] = User.Address.State
	}
	if User.Address.Country != "" {
		userToUpdate["address.country"] = User.Address.Country
	}

	if User.Hobbies != nil {
		userToUpdate["hobbies"] = User.Hobbies
	}

	update := bson.M{"$set": userToUpdate}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		response := httpresponse{
			Message: "failed to update user",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	if result.MatchedCount == 0 {
		response := httpresponse{
			Message: "user not found",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	w.WriteHeader(http.StatusOK)

	response := httpresponse{
		Message: "user updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}
