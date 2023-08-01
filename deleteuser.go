package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		message := "Method not allowed"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusMethodNotAllowed)
		return
	}
	params := r.URL.Query()
	userID := params.Get("id")
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		message := "Invalid user ID"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusMethodNotAllowed)
		return
	}

	client, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("testdb").Collection("users")

	filter := bson.M{"User_id": userIDInt}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}
