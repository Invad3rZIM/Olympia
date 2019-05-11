package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type EventCache struct {
	eventCache   map[string]*Event
	signingCount map[string]int
}

const hoursInADay = 1440
const dayStart = 480

func (c *EventCache) GetShifts(username string) []*Event {
	events := []*Event{}

	for _, v := range c.eventCache {
		s := v.CurrentGuards

		for _, t := range s {
			if t.Username == username {
				events = append(events, v)
				break
			}
		}
	}

	return events
}

func (c *EventCache) CanScheduleEvent(eventName string, day int, startTime int, duration int) bool {

	endTime := startTime + duration

	if startTime < dayStart {
		return false
	}

	if endTime > hoursInADay {
		return false
	}

	//medal ceremony thingy..

	e, _ := c.GetEvent(eventName)

	if e.NN != e.Name {
		competition, _ := c.GetEvent(e.NN)
		if competition.Day > day {
			return false
		}
		if competition.Day < day {
		} else if startTime < (competition.StartTime + competition.Duration) {
			return false
		}

	} else {
		medal, _ := c.GetEvent("(MC) - " + e.NN)

		if medal.Day < day {
			return false
		}

		if medal.Day > day {
		} else if startTime > medal.StartTime {
			return false
		}
	}

	for k, v := range c.eventCache {
		if strings.HasPrefix(v.Name, "(S) - ") {
			continue
		}

		if v.StartTime == 0 || v.Duration == 0 {
			continue
		}

		if k == eventName {
			continue
		}

		if v.Day != day {
			continue
		}

		vBuffer := v.StartTime + v.Duration + 15
		eBuffer := startTime + duration + 15

		if v.StartTime <= eBuffer && v.StartTime >= startTime {
			return false
		}

		if vBuffer <= eBuffer && vBuffer >= startTime {
			return false
		}

		if startTime <= vBuffer && startTime >= v.StartTime {
			return false
		}

		if eBuffer <= vBuffer && endTime >= v.StartTime {
			return false
		}
	}

	return true
}

func (c *EventCache) CanScheduleSigning(eventName string, username string, day int, startTime int, duration int, user *User) bool {

	endTime := startTime + duration

	if startTime < dayStart {
		return false
	}

	if endTime > hoursInADay {
		return false
	}

	for k, v := range user.EventsParticipating {

		if v.StartTime == 0 || v.Duration == 0 {
			continue
		}

		if k == eventName {
			continue
		}

		if v.Day != day {
			continue
		}

		vBuffer := v.StartTime + v.Duration + 15
		eBuffer := startTime + duration + 15

		if v.StartTime <= eBuffer && v.StartTime >= startTime {
			return false
		}

		if vBuffer <= eBuffer && vBuffer >= startTime {
			return false
		}

		if startTime <= vBuffer && startTime >= v.StartTime {
			return false
		}

		if eBuffer <= vBuffer && endTime >= v.StartTime {
			return false
		}
	}

	return true
}

func (c *EventCache) Schedule(event string, day int, startTime int, duration int) {
	e, _ := c.GetEvent(event)

	e.Day = day
	e.StartTime = startTime
	e.Duration = duration
	e.EndTime = startTime + duration
}

func (c *EventCache) CreateSigning(arena *Arena, user *User, day int, startTime int, duration int) {
	//finds first unique name
	i := 1
	var s string

	for {
		s = fmt.Sprintf("(S) - %s %s Signing [%d]", user.First, user.Last, i)
		if _, ok := c.eventCache[s]; !ok {
			break
		}

		i++
	}

	e := Event{
		Day:       day,
		StartTime: startTime,
		Duration:  duration,
		EndTime:   startTime + duration,
		Name:      s,
		NN:        s,
		Arena:     arena,
		Hosting:   user.Username,
	}

	c.AddEvent(&e)
	user.EventsParticipating[s] = &e

}

func (c *EventCache) DropSigning(u *User, event string) {
	delete(c.eventCache, event)
	delete(u.EventsParticipating, event)
}

func NewEventCache() *EventCache {
	return &EventCache{
		eventCache: make(map[string]*Event),
	}
}

//needs to get from foreign cache later
func (c *EventCache) GetEvent(name string) (*Event, error) {
	if a, exists := c.eventCache[name]; !exists {
		return nil, errors.New("event not found")
	} else {
		return a, nil
	}
}

func (c *EventCache) GetAllEvents() []*Event {
	events := []*Event{}
	keys := []string{}

	for k, _ := range c.eventCache {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		events = append(events, c.eventCache[v])
	}

	return events
}

func (c *EventCache) GetSelectedEvents(es map[string]int) []*Event {
	events := []*Event{}
	keys := []string{}

	for k, _ := range es {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		events = append(events, c.eventCache[v])
	}

	return events
}

func (c *EventCache) SetPublicPrice(name string, price float64) {
	event, err := c.GetEvent(name)

	if err != nil {
		return
	}

	event.PublicPrice = price

}

func (c *EventCache) SetStaffPrice(name string, price float64) {
	event, err := c.GetEvent(name)

	if err != nil {
		return
	}

	event.StaffPrice = price
}

func (c *EventCache) AddEvent(a *Event) error {
	if _, exists := c.eventCache[a.Name]; !exists {
		c.eventCache[a.Name] = a
	} else {
		return errors.New("name in cache")
	}
	return nil
}

func (c *EventCache) HasEvent(name string) bool {
	_, ok := c.eventCache[name]

	//check foreign database here

	return ok
}
