package main

import (
	"errors"
)

type ArenaCache struct {
	arenaCache map[string]*Arena
}

func NewArenaCache() *ArenaCache {
	return &ArenaCache{
		arenaCache: make(map[string]*Arena),
	}
}

//needs to get from foreign cache later
func (c *ArenaCache) GetArena(name string) (*Arena, error) {
	if a, exists := c.arenaCache[name]; !exists {
		return nil, errors.New("arena not found")
	} else {
		return a, nil
	}
}
func (c *ArenaCache) GetAllArenas() []*Arena {
	arenas := []*Arena{}

	for _, v := range c.arenaCache {
		arenas = append(arenas, v)
	}

	return arenas
}
func (c *ArenaCache) AddArena(a *Arena) error {
	if _, exists := c.arenaCache[a.Name]; !exists {
		c.arenaCache[a.Name] = a
	} else {
		return errors.New("name in cache")
	}
	return nil
}

func (c *ArenaCache) HasArena(name string) bool {
	_, ok := c.arenaCache[name]

	//check foreign database here

	return ok
}
