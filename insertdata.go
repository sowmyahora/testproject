package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	Country string `bson:"country" json:"country"`
}

type user struct {
	UserID  int64    `bson:"user_id" json:"user_id"`
	Name    string   `bson:"name" json:"name"`
	Phone   string   `bson:"phone" json:"phone"`
	Address address  `bson:"address" json:"address"`
	Hobbies []string `bson:"hobbies" json:"hobbies"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodPost:
		insertUser(w, r)
	default:
		message := "Method not allowed"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusMethodNotAllowed)
	}
}

func insertUser(w http.ResponseWriter, r *http.Request) {

	validate(r)
	var User user

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	countersCollection := client.Database("testdb").Collection("counters")

	filter := bson.M{"_id": "user_id"}

	update := bson.M{"$inc": bson.M{"seq": 1}}
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	result := countersCollection.FindOneAndUpdate(context.Background(), filter, update, opt)

	if result.Err() != nil {
		fmt.Println("Findoneandupdate")
		log.Println(result.Err())
	}

	var counterDoc struct {
		Seq int64 `bson:"seq"`
	}
	err = result.Decode(&counterDoc)
	if err != nil {
		message := err
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
	}

	User.UserID = counterDoc.Seq

	collection := client.Database("testdb").Collection("users")

	_, err = collection.InsertOne(context.TODO(), User)
	if err != nil {
		message := err
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
	}

	w.WriteHeader(http.StatusCreated)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User inserted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	userIDStr := queryParams.Get("user_id")
	if userIDStr == "" {
		message := "Missing 'user_id' query parameter"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		message := "Invalid 'user_id' query parameter"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusBadRequest)
		return
	}

	validate(r)

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("testdb").Collection("users")
	filter := bson.M{"user_id": userID}

	var userToUpdate map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&userToUpdate); err != nil {
		message := "Failed to parse request body"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusInternalServerError)
		return
	}

	delete(userToUpdate, "user_id")
	update := bson.M{"$set": userToUpdate}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		message := "Failed to update user"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		message := "user not found"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}
