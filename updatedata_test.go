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

func TestUpdateUser(t *testing.T) {
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		Name:  "Ansh Lokhande",
		Phone: "000000000",
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
		message := "Failed to marshal user object: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
	}

	req, err := http.NewRequest("PUT", "/users/update", bytes.NewBuffer(userJSON))
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
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
	client, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	User := user{
		Name:  "Rakesh Lokhande",
		Phone: "000000000",
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
		message := "Failed to marshal user object: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
	}

	req, err := http.NewRequest("PUT", "/users/update", bytes.NewBuffer(userJSON))
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
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
