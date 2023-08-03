package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userIDstr := mux.Vars(r)["user_id"]

	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		response := httpresponse{
			Message: "Invalid User ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	client, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("testdb").Collection("users")

	filter := bson.M{"user_id": userID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		response := httpresponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := httpresponse{
		Message: "user deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}
