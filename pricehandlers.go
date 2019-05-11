package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) EventSetStaffPrice(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "price"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	price := requestBody["price"].(float64)

	h.EventCache.SetStaffPrice(event, price)

	json.NewEncoder(w).Encode(&event)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) EventSetPublicPrice(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "price"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	event := requestBody["event"].(string)
	price := requestBody["price"].(float64)

	h.EventCache.SetPublicPrice(event, price)

	json.NewEncoder(w).Encode(&event)
	w.WriteHeader(http.StatusOK)
}
