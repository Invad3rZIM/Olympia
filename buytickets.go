package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "username", "count", "event", "free"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	name := requestBody["username"].(string)
	count := int(requestBody["count"].(float64))
	event := requestBody["event"].(string)
	free := requestBody["free"].(string)

	e, _ := h.GetEvent(event)
	u, _ := h.GetUser(name)

	if !(count <= e.RemainingTickets()) {
		return
	}

	e.TicketCount += count
	u.PurchaseTicket(e.Name, e.PublicPrice, e.StaffPrice, count, free == "on")

	json.NewEncoder(w).Encode(&MyEvents{Registered: u.Registrations})
	w.WriteHeader(http.StatusOK)
}
