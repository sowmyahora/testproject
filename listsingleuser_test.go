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

func TestListSingleUser(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Println(err)
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
			message := err
			jmsg, _ := json.Marshal(message)
			fmt.Println(jmsg)
		}
	}

	req, err := http.NewRequest("GET", "/users/890078", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
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

	req, err := http.NewRequest("GET", "/users/890078", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
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
