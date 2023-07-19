package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUpdateUser(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	user := User{
		user_id: 97659,
		name:    "Ansh Lokhande",
		phone:   "000000000",
		address: address{
			street:  "street 151",
			city:    "Jammu",
			state:   "J&K",
			country: "India",
		},
		hobbies: []string{"Reading", "Gaming", "Cooking"},
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user object: %v", err)
	}

	req, err := http.NewRequest("POST", "/users/update", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	updateUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := "User updated successfully"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}

func TestUpdateUserInvalidData(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	user := User{
		user_id: 2001,
		name:    "Rakesh Lokhande",
		phone:   "000000000",
		address: address{
			street:  "street 151",
			city:    "Jammu",
			state:   "J&K",
			country: "India",
		},
		hobbies: []string{"Reading", "Gaming", "Cooking"},
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user object: %v", err)
	}

	req, err := http.NewRequest("POST", "/users/update", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	updateUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := "User updated successfully"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}
