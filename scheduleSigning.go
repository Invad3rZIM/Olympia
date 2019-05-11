package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CanScheduleSigning(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "username", "day", "start", "duration"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	username := requestBody["username"].(string)
	day := int(requestBody["day"].(float64))
	start := int(requestBody["start"].(float64))
	duration := int(requestBody["duration"].(float64))

	user, _ := h.GetUser(username)

	b := Boolean{
		CanSchedule: h.EventCache.CanScheduleSigning(event, username, day, start, duration, user),
	}

	json.NewEncoder(w).Encode(&b)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateSigning(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "arena", "username", "day", "start", "duration"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	arena := requestBody["arena"].(string)
	username := requestBody["username"].(string)
	day := int(requestBody["day"].(float64))
	start := int(requestBody["start"].(float64))
	duration := int(requestBody["duration"].(float64))

	u, _ := h.GetUser(username)
	a, _ := h.GetArena(arena)

	h.EventCache.CreateSigning(a, u, day, start, duration)

	json.NewEncoder(w).Encode("4")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DropSigning(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	eventName := requestBody["event"].(string)
	username := requestBody["username"].(string)

	u, _ := h.GetUser(username)
	h.EventCache.DropSigning(u, eventName)

	json.NewEncoder(w).Encode("4")
	w.WriteHeader(http.StatusOK)
}
