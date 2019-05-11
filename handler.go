package main

import (
	"errors"
	"fmt"
)

type Handler struct {
	*UserCache
	*ArenaCache
	*EventCache
}


func NewHandler() *Handler {
	return &Handler{
		NewUserCache(),
		NewArenaCache(),
		NewEventCache(),
	}
}
/*
//authenticates the user via numeric pin... replace with JWT if time permits later
func (h *Handler) VerifyPin(uid int, pin int) error {
	u, err := h.GetUser(uid)

	//if user cannot be found for whatever reason
	if err != nil {
		return err
	}

	if !u.CheckPin(pin) {
		return errors.New("error: invalid pin")
	}

	return nil
}*/


//VerifyBody is a helper function to ensure all http requests contain the requisite fields returns error if fields missing
func (h *Handler) VerifyBody(body map[string]interface{}, str ...string) error {
	for _, s := range str {
		fmt.Println(s)
		if _, ok := body[s]; !ok {
			return errors.New("error: missing field: " + s)
		}
	}

	return nil
}
