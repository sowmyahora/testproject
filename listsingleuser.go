package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDstr := mux.Vars(r)["user_id"]

	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		message := "Invalid User ID"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusBadRequest)
		return
	}

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("testdb").Collection("users")

	filter := bson.M{"user_id": userID}

	var User user
	err = collection.FindOne(ctx, filter).Decode(&User)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			response := httpresponse{
				Message: "user does not exists",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := httpresponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return

	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		response := httpresponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
