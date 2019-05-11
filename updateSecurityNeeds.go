package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func (h *Handler) UpdateSecurityNeeds(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "event", "needed"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	event := requestBody["event"].(string)
	needed := int(requestBody["needed"].(float64))

	e, err := h.GetEvent(event)

	if err != nil {
		return
	}

	e.SecurityNeeded = needed
	e.GuardReset()

	//START
	//debug here

	guards := h.UserCache.GetSecurityGuards() //in the future, sort this list from least shifts to most

	Shuffle(guards)

	for _, v := range guards {
		if e.SecurityNeeded <= e.CurrentSecurity {
			break
		}

		if !e.OnDuty(v) {
			e.AddGuard(v)
		}
	}

	json.NewEncoder(w).Encode(&guards)
	w.WriteHeader(http.StatusOK)
}

func Shuffle(vals []*User) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}
