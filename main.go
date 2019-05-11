package main

import (
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

func main() {

	h := NewHandler()

	http.HandleFunc("/users/login", h.AuthenticateUser)
	http.HandleFunc("/users/taken", h.AccountExists)
	http.HandleFunc("/users/new", h.CreateUser)

	http.HandleFunc("/athlete/all", h.GetAllAthletes)
	http.HandleFunc("/athlete/schedule", h.SetEventAthlete)
	http.HandleFunc("/athlete/bio", h.SetAthleteBio)
	http.HandleFunc("/athlete/signing/open", h.CanScheduleSigning)
	http.HandleFunc("/athlete/signing", h.CreateSigning)
	http.HandleFunc("/athlete/signing/drop", h.DropSigning)

	http.HandleFunc("/arena/new", h.CreateArena)
	http.HandleFunc("/arena/taken", h.ArenaExists)
	http.HandleFunc("/arena/all", h.GetAllArenas)

	http.HandleFunc("/event/new", h.CreateEvent)
	http.HandleFunc("/event/taken", h.EventExists)
	http.HandleFunc("/event/mine", h.GetMyEvents)

	http.HandleFunc("/event/schedule", h.ScheduleEvent)
	http.HandleFunc("/event/schedule/open", h.CanScheduleEvent)

	http.HandleFunc("/event/price/public", h.EventSetPublicPrice)
	http.HandleFunc("/event/price/staff", h.EventSetStaffPrice)
	http.HandleFunc("/event/ticket", h.BuyTickets)

	http.HandleFunc("/event/pair", h.SetEventArena)
	http.HandleFunc("/event/all", h.GetAllEvents)

	http.HandleFunc("/security/update", h.UpdateSecurityNeeds)
	http.HandleFunc("/security/all", h.AllGuards)
	http.HandleFunc("/security/shifts", h.GetShifts)

	appengine.Main() // Starts the server to receive requests
}
