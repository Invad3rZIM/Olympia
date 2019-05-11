package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type UserCache struct {
	usercache map[string]*User
}

func NewUserCache() *UserCache {
	return &UserCache{
		usercache: make(map[string]*User),
	}
}

func (c *UserCache) GetMyEvents(username string) *MyEvents {
	u, _ := c.GetUser(username)

	return &MyEvents{Registered: u.Registrations}

}

func (c *UserCache) GetSecurityGuards() []*User {
	guards := []*User{}

	fmt.Printf(fmt.Sprint(len(c.usercache)))

	for _, v := range c.usercache {
		if v.UserType == "guard" {
			guards = append(guards, v)
		}
	}

	return guards
}

type MyEvents struct {
	Registered map[string]int
	MyEvents   []*Event
}

//needs to get from foreign cache later
func (c *UserCache) GetUser(username string) (*User, error) {
	if u, exists := c.usercache[username]; !exists {
		return nil, errors.New("user not found")
	} else {
		return u, nil
	}
}

//needs to get from foreign cache later
func (c *UserCache) Authenticate(username string, password string) (*User, error) {
	if u, exists := c.usercache[username]; !exists {
		return nil, errors.New("user not found")
	} else if u.Password != password {
		return nil, errors.New("invalid password")
	} else {
		return u, nil
	}
}

func (c *UserCache) AddUser(u *User) error {
	if _, exists := c.usercache[u.Username]; !exists {
		c.usercache[u.Username] = u
	} else {
		return errors.New("userid in cache")
	}
	return nil
}

func (c *UserCache) GetAllAthletes() []*User {
	athletes := []*User{}

	for _, v := range c.usercache {
		if v.UserType == "athlete" {
			athletes = append(athletes, v)
		}
	}

	return athletes
}

func (c *UserCache) HasUser(username string) bool {
	_, ok := c.usercache[username]

	//check foreign database here

	return ok
}

func (c *UserCache) GenSessionKey(u *User) {
	u.Key = rand.Intn(9999999)
}
