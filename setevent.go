package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) SetEventArena(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "arena"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	arena := requestBody["arena"].(string)

	a, err := h.GetArena(arena)

	if err != nil {

	}

	e, err := h.GetEvent(event)

	if err != nil {

	}

	e.Arena = a
	e.ArenaCapacity = a.Capacity

	json.NewEncoder(w).Encode(&e)
	w.WriteHeader(http.StatusOK)
}
