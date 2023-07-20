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

	User := user{
		User_id: 97659,
		Name:    "Ansh Lokhande",
		Phone:   "000000000",
		Address: address{
			Street:  "street 151",
			City:    "Jammu",
			State:   "J&K",
			Country: "India",
		},
		Hobbies: []string{"Reading", "Gaming", "Cooking"},
	}

	userJSON, err := json.Marshal(User)
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

	User := user{
		User_id: 2001,
		Name:    "Rakesh Lokhande",
		Phone:   "000000000",
		Address: address{
			Street:  "street 151",
			City:    "Jammu",
			State:   "J&K",
			Country: "India",
		},
		Hobbies: []string{"Reading", "Gaming", "Cooking"},
	}

	userJSON, err := json.Marshal(User)
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
