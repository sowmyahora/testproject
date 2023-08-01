package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		message := "Method not allowed"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" {
		message := "User ID not provided"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
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
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	filter := bson.M{"User_id": userID}

	var User user
	err = collection.FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		message := "Error retrieving user:"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err, http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		message := "Error marshaling user to JSON:"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
