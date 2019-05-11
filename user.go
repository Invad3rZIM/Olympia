package main

import "fmt"

type User struct {
	First    string
	Last     string
	Username string
	Password string
	UserType string

	Height  int
	Weight  int
	Country string
	Sex     string
	Age     int
	Bio     string

	Shifts int

	Registrations map[string]int
	Bill          float64

	EventsParticipating map[string]*Event

	Salt int
	Key  int
}

func NewUser(fn string, ln string, username string, pass string, usertype string) *User {
	u := User{
		First:               fn,
		Last:                ln,
		Username:            username,
		Password:            pass,
		UserType:            usertype,
		Registrations:       make(map[string]int),
		EventsParticipating: make(map[string]*Event),
	}

	return &u
}

func (u *User) AddEvent(e *Event) {
	fmt.Printf("%v\n\n\n\n%v\n\n", e, u.EventsParticipating)

	f := Event{
		Name:  e.Name,
		Arena: e.Arena,
	}
	u.EventsParticipating[e.Name] = &f
}

func (u *User) DropEvent(e *Event) {
	u.EventsParticipating[e.Name] = nil
	delete(u.EventsParticipating, e.Name)
}

func (u *User) GetEventsParticipating() []*Event {
	events := []*Event{}

	for _, v := range u.EventsParticipating {
		events = append(events, v)
	}

	return events
}

func (u *User) PurchaseTicket(eventName string, publicPrice float64, staffPrice float64, count int, free bool) {
	if u.UserType == "staff" {
		u.Bill += staffPrice * float64(count)
	} else if u.UserType == "athlete" {
		if !free {
			u.Bill += publicPrice * float64(count)
		}
	} else if u.UserType == "public" {
		u.Bill += publicPrice * float64(count)
	}

	u.Registrations[eventName] = u.Registrations[eventName] + count
}

//hashes and salts the password
func (u *User) HashAndSalt(pass string, salt int) {
	u.Salt = salt
	u.Password = pass //need to do the cryptography thingy later
}
