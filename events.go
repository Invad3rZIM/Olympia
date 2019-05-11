package main

type Event struct {
	Eid  int
	Name string

	Hosting string

	IsOutdoors bool
	IsAquatic  bool

	HasField bool
	HasTrack bool

	Day       int
	StartTime int
	EndTime   int
	Duration  int

	Arena         *Arena
	ArenaCapacity int

	IsMedal bool

	Athletes map[string]*User

	TicketCount int
	PublicPrice float64
	StaffPrice  float64

	SecurityNeeded  int
	CurrentSecurity int

	NN            string
	CurrentGuards []*User
}

func (e *Event) RemainingTickets() int {
	return e.ArenaCapacity - e.TicketCount
}

func (e *Event) AddGuard(u *User) {
	e.CurrentGuards = append(e.CurrentGuards, u)
	e.CurrentSecurity++
	u.Shifts++
}

func (e *Event) GetAthletesParticipating() []*User {
	athletes := []*User{}

	for _, v := range e.Athletes {
		athletes = append(athletes, v)
	}

	return athletes
}

func (e *Event) AddAthlete(u *User) {
	e.Athletes[u.Username] = u
}

func (e *Event) DropAthlete(u *User) {
	e.Athletes[u.Username] = nil
	delete(e.Athletes, u.Username)
}

func (e *Event) GetAthletes() []*User {
	list := []*User{}

	for _, v := range e.Athletes {
		if v != nil {
			list = append(list, v)
		}
	}

	return list
}

func (e *Event) GuardReset() {
	for _, v := range e.CurrentGuards {
		v.Shifts--
	}

	e.CurrentSecurity = 0
	e.CurrentGuards = []*User{}
}

func (e *Event) OnDuty(u *User) bool {
	for _, v := range e.CurrentGuards {
		if v == u {
			return true
		}
	}

	return false
}

func NewEvent(name string, duration int, isOutdoors bool, isAquatic bool, hasField bool, hasTrack bool, isMedal bool, nn string) *Event {
	a := Arena{}

	return &Event{
		Name:       name,
		IsOutdoors: isOutdoors,
		IsAquatic:  isAquatic,
		HasField:   hasField,
		HasTrack:   hasTrack,
		Duration:   duration,
		Arena:      &a,
		IsMedal:    isMedal,
		NN:         nn,
		Athletes:   make(map[string]*User),
	}
}
