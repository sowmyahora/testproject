package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestDeleteUser(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/users/5", nil)
	//req, err := http.NewRequest("DELETE", "mongodb://mongo:27017/users/delete?id=890078", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
	}

	vars := map[string]string{
		"user_id": "5",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	deleteUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestDeleteUserInvalidData(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/users/5", nil)
	//req, err := http.NewRequest("DELETE", "mongodb://mongo:27017/users/delete?id=890078", nil)
	if err != nil {
		message := "Failed to create request: %v"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg, err)
	}

	vars := map[string]string{
		"user_id": "5",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	deleteUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
