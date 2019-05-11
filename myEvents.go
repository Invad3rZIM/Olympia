package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetMyEvents(w http.ResponseWriter, r *http.Request) {
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

	myEvents := h.UserCache.GetMyEvents(username)

	myEvents.MyEvents = h.EventCache.GetSelectedEvents(myEvents.Registered)

	json.NewEncoder(w).Encode(&myEvents)
	w.WriteHeader(http.StatusOK)
}
