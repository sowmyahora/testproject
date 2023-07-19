package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestListUsers(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	users := []User{
		{
			user_id: 890078,
			name:    "Sonia Khera",
			phone:   "9910470030",
			address: address{
				street:  "Street 1",
				city:    "New York",
				state:   "NY",
				country: "USA",
			},
			hobbies: []string{"Reading", "Gaming", "Cooking"},
		},
		{
			user_id: 817886,
			name:    "Sam Manchanda",
			phone:   "987657899",
			address: address{
				street:  "Street 2",
				city:    "Los Angeles",
				state:   "CA",
				country: "USA",
			},
			hobbies: []string{"Traveling", "Photography", "Painting"},
		},
	}

	collection := client.Database("testdb").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}
	}

	req, err := http.NewRequest("GET", "/users/list", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	listUsers(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := `[{"User_id":890078,"Name":"Sonia Khera","Phone":"9910470030","Address":{"Street":"Street 1","City":"New York","State":"NY","Country":"USA"},"Hobbies":["Reading","Gaming","Cooking"]},{"User_id":817886,"Name":"Sam Manchanda","Phone":"987657899","Address":{"Street":"Street 2","City":"Los Angeles","State":"CA","Country":"USA"},"Hobbies":["Traveling","Photography","Painting"]}]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}

func TestListUsersInvalidData(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	users := []User{
		{

			user_id: 890078,
			name:    "Sonia Khera",
			phone:   "9910470030",
			address: address{
				street:  "Street 1",
				city:    "New York",
				state:   "NY",
				country: "USA",
			},
			hobbies: []string{"Reading", "Gaming", "Cooking"},
		},
		{
			user_id: 817886,
			name:    "Sam Manchanda",
			phone:   "987657899",
			address: address{
				street:  "Street 2",
				city:    "Los Angeles",
				state:   "CA",
				country: "USA",
			},
			hobbies: []string{"Traveling", "Photography", "Painting"},
		},
	}

	collection := client.Database("testdb").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}
	}

	req, err := http.NewRequest("GET", "/users/list", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	listUsers(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedResponse := `[{"User_id":90121,"Name":"John Doe","Phone":"1234567890","Address":{"Street":"Street 1","City":"New York","State":"NY","Country":"USA"},"Hobbies":["Reading","Gaming","Cooking"]},{"User_id":90122,"Name":"Jane Smith","Phone":"9876543210","Address":{"Street":"Street 2","City":"Los Angeles","State":"CA","Country":"USA"},"Hobbies":["Traveling","Photography","Painting"]}]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %q, got %q", expectedResponse, rr.Body.String())
	}
}
