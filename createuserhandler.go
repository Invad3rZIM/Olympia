package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "firstname", "lastname", "username", "password", "usertype"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	fn := requestBody["firstname"].(string)
	ln := requestBody["lastname"].(string)
	username := requestBody["username"].(string)
	pass := requestBody["password"].(string)
	usertype := requestBody["usertype"].(string)

	u := NewUser(fn, ln, username, pass, usertype)

	h.AddUser(u)
	h.UserCache.GenSessionKey(u)

	json.NewEncoder(w).Encode(&u)
	w.WriteHeader(http.StatusOK)
}
