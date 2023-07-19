package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestListSingleUser(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	users := []User{
		{
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
		},
		{
			User_id: 817886,
			Name:    "Sam Manchanda",
			Phone:   "987657899",
			Address: address{
				Street:  "Street 2",
				City:    "Los Angeles",
				State:   "CA",
				Country: "USA",
			},
			Hobbies: []string{"Traveling", "Photography", "Painting"},
		},
	}

	collection := client.Database("testdb").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}
	}

	req, err := http.NewRequest("GET", "/users/890078", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	getUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := `{"User_id":890078,"Name":"Sonia Khera","Phone":"9910470030","Address":{"Street":"Street 1","City":"New York","State":"NY","Country":"USA"},"Hobbies":["Reading","Gaming","Cooking"]}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}

func TestListSingleUserInvalidData(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	users := []User{
		{
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
		},
		{
			User_id: 817886,
			Name:    "Sam Manchanda",
			Phone:   "987657899",
			Address: address{
				Street:  "Street 2",
				City:    "Los Angeles",
				State:   "CA",
				Country: "USA",
			},
			Hobbies: []string{"Traveling", "Photography", "Painting"},
		},
	}

	collection := client.Database("testdb").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}
	}

	req, err := http.NewRequest("GET", "/users/890078", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	getUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := `{"User_id":79008,"Name":"Sonia Khera","Phone":"9910470030","Address":{"Street":"Street 1","City":"New York","State":"NY","Country":"USA"},"Hobbies":["Reading","Gaming","Cooking"]}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}