package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsers(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	Users := []user{
		{
			Name:  "Sonia Khera",
			Phone: "9910470030",
			Address: address{
				Street:  "Street 1",
				City:    "New York",
				State:   "NY",
				Country: "USA",
			},
			Hobbies: []string{"Reading", "Gaming", "Cooking"},
		},
		{
			Name:  "Sam Manchanda",
			Phone: "987657899",
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
	for _, user := range Users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			message := "Failed to insert user: %v"
			jmsg, _ := json.Marshal(message)
			fmt.Println(jmsg, err)
		}
	}

	req, err := http.NewRequest("GET", "/users/list", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
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
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	Users := []user{
		{
			Name:  "Sonia Khera",
			Phone: "9910470030",
			Address: address{
				Street:  "Street 1",
				City:    "New York",
				State:   "NY",
				Country: "USA",
			},
			Hobbies: []string{"Reading", "Gaming", "Cooking"},
		},
		{
			Name:  "Sam Manchanda",
			Phone: "987657899",
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
	for _, user := range Users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			message := "Failed to insert user: %v"
			jmsg, _ := json.Marshal(message)
			fmt.Println(jmsg, err)
		}
	}

	req, err := http.NewRequest("GET", "/users/list", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
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
