package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateArena(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "name", "cap", "outdoors", "aquatic", "field", "track"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// //extract field values to variables for readability
	name := requestBody["name"].(string)
	cap := int(requestBody["cap"].(float64))
	outdoors := requestBody["outdoors"].(string)
	aquatic := requestBody["aquatic"].(string)
	field := requestBody["field"].(string)
	track := requestBody["track"].(string)

	fmt.Printf("%s %d %s ", name, cap, outdoors)

	a := NewArena(name, cap, outdoors == "on", aquatic == "on", field == "on", track == "on")

	h.AddArena(a)

	json.NewEncoder(w).Encode(&a)
	w.WriteHeader(http.StatusOK)
}
