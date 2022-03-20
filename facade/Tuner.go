package main

import (
	"fmt"
	"strconv"
)

type Tuner struct {
	description string
	frequency   float32
}

func NewTuner(description string) *Tuner {
	return &Tuner{description: description}
}

func (t *Tuner) On() {
	fmt.Println(t.description + " On")
}

func (t *Tuner) Off() {
	fmt.Println(t.description + " Off")
}

func (t *Tuner) SetFrequency(frequency float32) {
	fmt.Println(t.description + " setting frequency to " + strconv.FormatFloat(float64(frequency), 'f', -1, 32))
	t.frequency = frequency
}

func (t *Tuner) SetAM() {
	fmt.Println(t.description + " setting in AM mode")
}

func (t *Tuner) SetFM() {
	fmt.Println(t.description + " setting in FM mode")
}

func (t *Tuner) String() string {
	return t.description
}
