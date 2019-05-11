package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) SetAthleteBio(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "username"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	username := requestBody["username"].(string)

	u, _ := h.GetUser(username)

	if weight, ok := requestBody["weight"].(float64); ok {
		if weight > -1 {
			u.Weight = int(weight)
		}
	}

	if height, ok := requestBody["height"].(float64); ok {
		if height > -1 {
			u.Height = int(height)
		}
	}

	if age, ok := requestBody["age"].(float64); ok {
		if age > -1 {
			u.Age = int(age)
		}
	}

	if bio, ok := requestBody["bio"].(string); ok {
		if bio != "" {
			u.Bio = bio
		}
	}

	if country, ok := requestBody["country"].(string); ok {
		if country != "" {
			u.Country = country
		}
	}

	if sex, ok := requestBody["sex"].(string); ok {
		if sex != "" {
			u.Sex = sex
		}
	}

	json.NewEncoder(w).Encode(&u)
	w.WriteHeader(http.StatusOK)
}
