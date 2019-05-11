package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAllAthletes(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	athletes := AllAthletes{h.UserCache.GetAllAthletes()}

	json.NewEncoder(w).Encode(&athletes)
	w.WriteHeader(http.StatusOK)
}

type AllAthletes struct {
	AllAthletes []*User
}
