package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func listUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		message := "Method not allowed"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusMethodNotAllowed)
		return
	}
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		message := ("Error retrieving users:")
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var Users []user
	for cursor.Next(context.TODO()) {
		var User user
		err := cursor.Decode(&User)
		if err != nil {
			message := "error"
			jmsg, _ := json.Marshal(message)
			fmt.Println(jmsg)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		Users = append(Users, User)
	}

	usersJSON, err := json.Marshal(Users)
	if err != nil {
		message := "Error marshaling users to JSON:"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}
