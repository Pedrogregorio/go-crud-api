package main

import (
	"encoding/json"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, status int, users []User, user User, errors []string) {
	if status != 200 {
		responseError(w, status, errors)
		return
	}

	hasUsers := len(users) > 0

	if hasUsers {
		responseUsers(w, status, users)
		return
	}

	responseUser(w, status, user)
}

func responseError(w http.ResponseWriter, status int, errors []string) {
	response := BaseResposeWithError{Status: status, Errors: errors}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func responseUser(w http.ResponseWriter, status int, user User) {
	response := BaseResposeWithUser{Status: status, Response: user}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func responseUsers(w http.ResponseWriter, status int, users []User) {
	response := BaseResposeWithUsers{Status: status, Response: users}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
