package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestListSingleUser(t *testing.T) {

	req, err := http.NewRequest("GET", "/users/0", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
	}

	vars := map[string]string{
		"user_id": "0",
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()

	getUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestListSingleUserInvalidData(t *testing.T) {

	req, err := http.NewRequest("GET", "/users/abcd", nil)
	if err != nil {

		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
	}

	rr := httptest.NewRecorder()

	getUser(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
