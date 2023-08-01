package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func validate(r *http.Request) {

	var User user

	err := json.NewDecoder(r.Body).Decode(&User)

	if err != nil {
		message := "Failed to parse request body"
		jmsg, _ := json.Marshal(message)
		fmt.Println(jmsg)
		return

	}

	if User.Name == "" {
		message := "Invalid user: Name cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if User.Phone == "" {
		message := "Invalid user: Phone cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if !isNumeric(User.Phone) {
		message := "Invalid user: phone can only be numeric"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if User.Address.Street == "" {
		message := "Invalid user: Street cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if User.Address.City == "" {
		message := "Invalid user: City cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if User.Address.State == "" {
		message := "Invalid user: State cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if User.Address.Country == "" {
		message := "Invalid user: Country cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	if len(User.Hobbies) == 0 {
		message := "Invalid user: Hobbies cannot be empty"
		jsmg, _ := json.Marshal(message)
		fmt.Println(jsmg, http.StatusBadRequest)
		return
	}

	for _, hobby := range User.Hobbies {
		if hobby == "" {
			message := "Invalid user: Hobbies cannot be empty"
			jsmg, _ := json.Marshal(message)
			fmt.Println(jsmg, http.StatusBadRequest)
			return
		}
	}
}
