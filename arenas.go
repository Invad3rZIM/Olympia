package main

type Arena struct {
	Aid  int
	Name string

	Capacity int

	IsOutdoors bool
	IsAquatic  bool

	HasField bool
	HasTrack bool
}

func NewArena(name string, cap int, isOutdoors bool, isAquatic bool, hasField bool, hasTrack bool) *Arena {
	return &Arena{
		Name:       name,
		Capacity:   cap,
		IsOutdoors: isOutdoors,
		IsAquatic:  isAquatic,
		HasField:   hasField,
		HasTrack:   hasTrack,
	}
}
