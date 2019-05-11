package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) AllGuards(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	guards := h.UserCache.GetSecurityGuards() //in the future, sort this list from least shifts to most

	//e, needed) //updates the security guards who'se gonna be taking on the shifts

	json.NewEncoder(w).Encode(&guards)
	w.WriteHeader(http.StatusOK)
}
