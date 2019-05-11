package main

import "math/rand"

func GenSalt() int {
	return rand.Intn(999) + 4
}

func GenUID() int {
	return 6
}
