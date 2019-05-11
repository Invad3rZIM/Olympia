package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAllArenas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	arenas := AllArenas{h.ArenaCache.GetAllArenas()}

	json.NewEncoder(w).Encode(&arenas)
	w.WriteHeader(http.StatusOK)
}

type AllArenas struct {
	AllArenas []*Arena
}
