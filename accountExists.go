package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h Handler) AccountExists(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "username"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	username := requestBody["username"].(string)

	if _, err := h.GetUser(username); err != nil {

		json.NewEncoder(w).Encode(&Status{Available: true})
		w.WriteHeader(http.StatusOK)
	} else {
		json.NewEncoder(w).Encode(&Status{Available: false})
		w.WriteHeader(http.StatusOK)
	}
}

type Status struct {
	Available bool
}
