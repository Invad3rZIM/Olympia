package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CanScheduleEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "day", "start", "duration"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	day := int(requestBody["day"].(float64))
	start := int(requestBody["start"].(float64))
	duration := int(requestBody["duration"].(float64))

	b := Boolean{
		CanSchedule: h.EventCache.CanScheduleEvent(event, day, start, duration),
	}
	json.NewEncoder(w).Encode(&b)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ScheduleEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "day", "start", "duration"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	day := int(requestBody["day"].(float64))
	start := int(requestBody["start"].(float64))
	duration := int(requestBody["duration"].(float64))

	h.EventCache.Schedule(event, day, start, duration)

	json.NewEncoder(w).Encode("4")
	w.WriteHeader(http.StatusOK)
}

type Boolean struct {
	CanSchedule bool
}
