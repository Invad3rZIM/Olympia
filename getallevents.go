package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	events := AllEvents{h.EventCache.GetAllEvents()}

	json.NewEncoder(w).Encode(&events)
	w.WriteHeader(http.StatusOK)
}

type AllEvents struct {
	AllEvents []*Event
}
