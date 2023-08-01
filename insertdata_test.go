package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertUser(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		Name:  "Ansh Tiwari",
		Phone: "8178860317",
		Address: address{
			Street:  "street 27",
			City:    "Pune",
			State:   "Maharashtra",
			Country: "India",
		},
		Hobbies: []string{"Playing Cricket", "Cooking", "Swimming"},
	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		message := "Failed to marshal user object: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return
	}

	req, err := http.NewRequest("POST", "/users/insert", bytes.NewBuffer(userJSON))
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return
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

func TestInsertUserInvalidData(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		Name:  "",
		Phone: "8178860317",
		Address: address{
			Street:  "street 27",
			City:    "Pune",
			State:   "Maharashtra",
			Country: "India",
		},
		Hobbies: []string{"Playing Cricket", "Cooking", "Swimming"},
	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		message := "Failed to marshal user object: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return
	}

	req, err := http.NewRequest("POST", "/users/insert", bytes.NewBuffer(userJSON))
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return
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
