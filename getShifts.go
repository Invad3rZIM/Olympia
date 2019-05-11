package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetShifts(w http.ResponseWriter, r *http.Request) {
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

	events := h.EventCache.GetShifts(username) //in the future, sort this list from least shifts to most

	json.NewEncoder(w).Encode(&AllEvents{events})
	w.WriteHeader(http.StatusOK)
}
