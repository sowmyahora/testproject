package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		User_id: 890078,
		Name:    "Sonia Khera",
		Phone:   "9910470030",
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
		t.Fatalf("Failed to insert user: %v", err)
	}

	//req, err := http.NewRequest("DELETE", "http://localhost:8080/users/delete?id=890078", nil)
	req, err := http.NewRequest("DELETE", "mongodb://mongo:27017/users/delete?id=890078", nil)
	if err != nil {
		t.Fatalf("Failed to create delete request: %v", err)
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
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		User_id: 877,
		Name:    "Sonia Khera",
		Phone:   "9910470030",
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
		t.Fatalf("Failed to insert user: %v", err)
	}

	//req, err := http.NewRequest("DELETE", "http://localhost:8080/users/delete?id=890", nil)
	req, err := http.NewRequest("DELETE", "mongodb://mongo:27017/users/delete?id=890078", nil)
	if err != nil {
		t.Fatalf("Failed to create delete request: %v", err)
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
