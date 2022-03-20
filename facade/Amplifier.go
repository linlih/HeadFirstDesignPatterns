package main

import (
	"fmt"
	"strconv"
)

type Amplifier struct {
	description string
	cd          CdPlayer
	dvd         DvdPlayer
	tuner       Tuner
}

func NewAmplifier(description string) *Amplifier {
	return &Amplifier{description: description}
}

func (a *Amplifier) On() {
	fmt.Println(a.description + " On")
}

func (a *Amplifier) Off() {
	fmt.Println(a.description + " Off")
}

func (a *Amplifier) SetStereoSound() {
	fmt.Println(a.description + " stereo mode on")
}

func (a *Amplifier) SetSurroundSound() {
	fmt.Println(a.description + " surround sound on (5 speakers, 1 subwoofer)")
}

func (a *Amplifier) SetVolume(level int) {
	fmt.Println(a.description + " setting volume level to " + strconv.FormatInt(int64(level), 10))
}

func (a *Amplifier) SetTuner(tuner Tuner) {
	a.tuner = tuner
	fmt.Println(a.description + " setting tuner to" + tuner.String())
}

func (a *Amplifier) SetDvd(dvd DvdPlayer) {
	a.dvd = dvd
	fmt.Println(a.description + " setting DVD player to" + dvd.String())
}

func (a *Amplifier) SetCd(cd CdPlayer) {
	a.cd = cd
	fmt.Println(a.description + " setting CD player to" + cd.String())
}

func (a *Amplifier) String() string {
	return a.description
}
