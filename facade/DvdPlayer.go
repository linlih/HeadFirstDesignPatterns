package main

import (
	"fmt"
	"strconv"
)

type DvdPlayer struct {
	description  string
	movie        string
	currentTrack int
}

func NewDvdPlayer(description string) *DvdPlayer {
	return &DvdPlayer{description: description}
}

func (d *DvdPlayer) On() {
	fmt.Println(d.description + " On")
}

func (d *DvdPlayer) Off() {
	fmt.Println(d.description + " Off")
}

func (d *DvdPlayer) Eject() {
	d.movie = ""
	fmt.Println(d.description + " eject")
}

func (d *DvdPlayer) Play(movie string) {
	d.movie = movie
	d.currentTrack = 0
	fmt.Println(d.description + " playing \"" + movie + "\"")
}

func (d *DvdPlayer) PlayByTrack(track int) {
	if d.movie == "" {
		fmt.Println(d.description + " can't play track " + strconv.FormatInt(int64(d.currentTrack), 10) + ", not dvd inserted")
	} else {
		d.currentTrack = track
		fmt.Println(d.description + " playing track " + strconv.FormatInt(int64(d.currentTrack), 10) + " of\"" + d.movie + "\"")
	}
}

func (d *DvdPlayer) Stop() {
	d.currentTrack = 0
	fmt.Println(d.description + " stopped \"" + d.movie + "\"")
}

func (d *DvdPlayer) Paused() {
	d.currentTrack = 0
	fmt.Println(d.description + " paused \"" + d.movie + "\"")
}

func (d *DvdPlayer) SetTwoChannelAudio() {
	fmt.Println(d.description + " setting two channel audio")
}

func (d *DvdPlayer) SetSurroundAudio() {
	fmt.Println(d.description + " setting surround audio")
}

func (d *DvdPlayer) String() string {
	return d.description
}
