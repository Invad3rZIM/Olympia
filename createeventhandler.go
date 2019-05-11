package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "name", "duration", "outdoors", "aquatic", "field", "track"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	name := requestBody["name"].(string)
	duration := int(requestBody["duration"].(float64))
	outdoors := requestBody["outdoors"].(string)
	aquatic := requestBody["aquatic"].(string)
	field := requestBody["field"].(string)
	track := requestBody["track"].(string)

	e := NewEvent(name, duration, outdoors == "on", aquatic == "on", field == "on", track == "on", true, name)
	m := NewEvent("(MC) - "+name, duration, outdoors == "on", aquatic == "on", field == "on", track == "on", false, name)

	h.EventCache.AddEvent(m)
	h.EventCache.AddEvent(e)

	json.NewEncoder(w).Encode(&e)
	w.WriteHeader(http.StatusOK)
}
