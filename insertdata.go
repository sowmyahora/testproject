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
	Street  string `bson:"Street"`
	City    string `bson:"City"`
	State   string `bson:"State"`
	Country string `bson:"Country"`
}

type user struct {
	User_id int64    `bson:"User_id"`
	Name    string   `bson:"Name"`
	Phone   string   `bson:"Phone"`
	Address address  `bson:"Address"`
	Hobbies []string `bson:"Hobbies"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodPost:
		insertUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func insertUser(w http.ResponseWriter, r *http.Request) {

	var User user

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if User.Name == "" {
		http.Error(w, "Invalid user: Name cannot be empty", http.StatusBadRequest)
		return
	}

	if User.Phone == "" {
		http.Error(w, "Invalid user: Phone cannot be empty", http.StatusBadRequest)
		return
	}

	if !isNumeric(User.Phone) {
		http.Error(w, "Invalid user: Phone number must contain only numeric digits", http.StatusBadRequest)
		return
	}

	if User.Address.Street == "" {
		http.Error(w, "Invalid user: Street in address cannot be empty", http.StatusBadRequest)
		return
	}

	if User.Address.City == "" {
		http.Error(w, "Invalid user: City in address cannot be empty", http.StatusBadRequest)
		return
	}

	if User.Address.State == "" {
		http.Error(w, "Invalid user: State in address cannot be empty", http.StatusBadRequest)
		return
	}

	if User.Address.Country == "" {
		http.Error(w, "Invalid user: Country in address cannot be empty", http.StatusBadRequest)
		return
	}

	if len(User.Hobbies) == 0 {
		http.Error(w, "Invalid user: At least one hobby is required", http.StatusBadRequest)
		return
	}

	for _, hobby := range User.Hobbies {
		if hobby == "" {
			http.Error(w, "Invalid user: Hobbies cannot have empty values", http.StatusBadRequest)
			return
		}
	}

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	countersCollection := client.Database("testdb").Collection("counters")

	filter := bson.M{"_id": "user_id"}

	update := bson.M{"$inc": bson.M{"seq": 1}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := countersCollection.FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		fmt.Println("Findoneandupdate")
		log.Println(result.Err())
	}

	var counterDoc struct {
		Seq int64 `bson:"seq"`
	}
	err = result.Decode(&counterDoc)
	if err != nil {
		log.Println(err)
	}

	User.User_id = counterDoc.Seq

	collection := client.Database("testdb").Collection("users")

	_, err = collection.InsertOne(context.TODO(), User)
	if err != nil {
		log.Println(err)
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
		http.Error(w, "Missing 'user_id' query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid 'user_id' query parameter", http.StatusBadRequest)
		return
	}

	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("testdb").Collection("users")
	filter := bson.M{"User_id": userID}

	var existingUser user
	err = collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		log.Println("Error retrieving user:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Printf("Existing user: %+v", existingUser)

	var User user
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	update := bson.M{"$set": bson.M{
		"Name":    User.Name,
		"Phone":   User.Phone,
		"Address": User.Address,
		"Hobbies": User.Hobbies,
	}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if result.ModifiedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
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
