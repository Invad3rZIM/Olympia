package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) SetEventAthlete(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "username", "toggle"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	username := requestBody["username"].(string)
	toggle := requestBody["toggle"].(string)

	e, _ := h.GetEvent(event)
	u, _ := h.GetUser(username)

	if toggle == "on" {
		e.AddAthlete(u)
		u.AddEvent(e)
	} else {
		e.DropAthlete(u)
		u.DropEvent(e)
	}

	json.NewEncoder(w).Encode(&e)
	w.WriteHeader(http.StatusOK)
}
