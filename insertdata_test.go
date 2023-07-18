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

func TestInsertuser(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	user := User{
		User_id: 90121,
		Name:    "Ravi Tiwari",
		Phone:   "8178860317",
		Address: address{
			Street:  "street 27",
			City:    "Pune",
			State:   "Maharashtra",
			Country: "India",
		},
		Hobbies: []string{"Playing Cricket", "Cooking", "Swimming"},
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user object: %v", err)
	}

	req, err := http.NewRequest("POST", "/users/insert", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	insertUser(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	expectedResponse := "User created successfully"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}

}

func TestInsertuserwrong(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	user := User{
		User_id: 90121,
		Name:    "",
		Phone:   "8178860317",
		Address: address{
			Street:  "street 27",
			City:    "Pune",
			State:   "Maharashtra",
			Country: "India",
		},
		Hobbies: []string{"Playing Cricket", "Cooking", "Swimming"},
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user object: %v", err)
	}

	req, err := http.NewRequest("POST", "/users/insert", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	insertUser(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}

	expectedResponse := "Invalid user: Name cannot be empty"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}
