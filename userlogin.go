package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (h Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "username", "password"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	username := requestBody["username"].(string)
	password := requestBody["password"].(string)

	if u, err := h.UserCache.Authenticate(username, password); err == nil {
		a := *u
		a.Password = ""
		json.NewEncoder(w).Encode(&a)
		w.WriteHeader(http.StatusOK)
	} else {
		json.NewEncoder(w).Encode(&User{})
		w.WriteHeader(http.StatusOK)
	}
}

type SessionKey struct {
	Key      int
	UserType string
}
