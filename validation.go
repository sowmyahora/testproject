package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func validateInsertRequest(r *http.Request) (user, error) {

	var User user

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return user{}, errors.New("Failed to parse request body")

	}

	err = json.Unmarshal(body, &User)

	if err != nil {
		message := "Failed to parse request body"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return user{}, errors.New("Failed to parse request body")

	}

	if User.Name == "" {
		return user{}, errors.New("Invalid user: Name cannot be empty")
	}

	if User.Phone == "" {
		return user{}, errors.New("Invalid user: Phone cannot be empty")
	}

	if !isNumeric(User.Phone) {
		return user{}, errors.New("Invalid user: phone can only be numeric")
	}

	if User.Address.Street == "" {
		return user{}, errors.New("Invalid user: Street cannot be empty")
	}

	if User.Address.City == "" {
		return user{}, errors.New("Invalid user: City cannot be empty")
	}

	if User.Address.State == "" {
		return user{}, errors.New("Invalid user: State cannot be empty")
	}

	if User.Address.Country == "" {
		return user{}, errors.New("Invalid user: Country cannot be empty")
	}

	if len(User.Hobbies) == 0 {
		return user{}, errors.New("Invalid user: Hobbies cannot be empty")
	}

	for _, hobby := range User.Hobbies {
		if hobby == "" {
			return user{}, errors.New("Invalid user: Hobbies cannot be empty")
		}

	}

	fmt.Println("Validation of fields successful")
	return User, nil

}

func validateUpdateRequest(r *http.Request) (user, error) {

	var User user

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return user{}, errors.New("Failed to parse request body")

	}

	err = json.Unmarshal(body, &User)

	if err != nil {
		message := "Failed to parse request body"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return user{}, errors.New("Failed to parse request body")

	}

	if !isNumeric(User.Phone) {
		message := "Invalid user: phone can only be numeric"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return user{}, errors.New("Invalid user: phone can only be numeric")
	}

	fmt.Println("Validation of fields successful")
	return User, nil

}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
