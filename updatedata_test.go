package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestUpdateUser(t *testing.T) {

	User := user{
		Name:  "Ansh Lokhande",
		Phone: "9887766644",
	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		message := "Failed to marshal user object: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
	}

	req, err := http.NewRequest("PUT", "/users/0", bytes.NewBuffer(userJSON))
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
	}

	vars := map[string]string{
		"user_id": "0",
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()

	updateUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateUserInvalidData(t *testing.T) {

	User := user{
		Name:  "Rakesh Lokhande",
		Phone: "000000000",
	}

	userJSON, err := json.Marshal(User)
	if err != nil {
		message := "Failed to marshal user object: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
	}

	req, err := http.NewRequest("PUT", "/users/0", bytes.NewBuffer(userJSON))
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
		return
	}

	vars := map[string]string{
		"user_id": "0",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()

	updateUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
