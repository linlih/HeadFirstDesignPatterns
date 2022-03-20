package main

import (
	"fmt"
	"strconv"
)

type TheaterLights struct {
	description string
}

func NewTheaterLights(description string) *TheaterLights {
	return &TheaterLights{description: description}
}

func (t *TheaterLights) On() {
	fmt.Println(t.description + " On")
}

func (t *TheaterLights) Off() {
	fmt.Println(t.description + " Off")
}

func (t *TheaterLights) Dim(level int) {
	fmt.Println(t.description + " dimming to " + strconv.FormatInt(int64(level), 10) + " %")
}

func (t *TheaterLights) String() string {
	return t.description
}
