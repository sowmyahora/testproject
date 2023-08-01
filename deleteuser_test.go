package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		Name:  "Sonia Khera",
		Phone: "9910470030",
		Address: address{
			Street:  "Street 1",
			City:    "New York",
			State:   "NY",
			Country: "USA",
		},
		Hobbies: []string{"Reading", "Gaming", "Cooking"},
	}

	collection := client.Database("testdb").Collection("users")
	_, err = collection.InsertOne(context.TODO(), User)
	if err != nil {

	}

	//req, err := http.NewRequest("DELETE", "http://localhost:8080/users/delete?id=890078", nil)
	req, err := http.NewRequest("DELETE", "mongodb://mongo:27017/users/delete?id=890078", nil)
	if err != nil {
		message := err
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
	}

	rr := httptest.NewRecorder()
	deleteUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Failed to delete user: Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := `{"message":"User deleted successfully"}`
	actualResponse := strings.TrimSpace(rr.Body.String())
	if actualResponse != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, actualResponse)
	}
}

func TestDeleteUserInvalidData(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		Name:  "Sonia Khera",
		Phone: "9910470030",
		Address: address{
			Street:  "Street 1",
			City:    "New York",
			State:   "NY",
			Country: "USA",
		},
		Hobbies: []string{"Reading", "Gaming", "Cooking"},
	}

	collection := client.Database("testdb").Collection("users")
	_, err = collection.InsertOne(context.TODO(), User)
	if err != nil {
		message := "Failed to insert user: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
	}

	//req, err := http.NewRequest("DELETE", "http://localhost:8080/users/delete?id=890", nil)
	req, err := http.NewRequest("DELETE", "mongodb://mongo:27017/users/delete?id=890078", nil)
	if err != nil {
		message := "Failed to create delete request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
	}

	rr := httptest.NewRecorder()
	deleteUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Failed to delete user: Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := `{"message":"User deleted successfully"}`
	actualResponse := strings.TrimSpace(rr.Body.String())
	if actualResponse != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, actualResponse)
	}
}
