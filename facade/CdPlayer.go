package main

import (
	"fmt"
	"strconv"
)

type CdPlayer struct {
	description  string
	title        string
	currentTrack int
}

func NewCdPlayer(description string) *CdPlayer {
	return &CdPlayer{description: description}
}

func (c *CdPlayer) On() {
	fmt.Println(c.description + " On")
}

func (c *CdPlayer) Off() {
	fmt.Println(c.description + " Off")
}

func (c *CdPlayer) Eject() {
	c.title = ""
	fmt.Println(c.description + " eject")
}

func (c *CdPlayer) Play(title string) {
	c.title = title
	c.currentTrack = 0
	fmt.Println(c.description + " playing \"" + title + "\"")
}

func (c *CdPlayer) PlayByTrack(track int) {
	if c.title == "" {
		fmt.Println(c.description + " can't play track " + strconv.FormatInt(int64(c.currentTrack), 10) + ", not cd inserted")
	} else {
		c.currentTrack = track
		fmt.Println(c.description + " playing track " + strconv.FormatInt(int64(c.currentTrack), 10))
	}
}

func (c *CdPlayer) Stop() {
	c.currentTrack = 0
	fmt.Println(c.description + " stopped")
}

func (c *CdPlayer) Pause() {
	fmt.Println(c.description + " paused \"" + c.title + "\"")
}

func (c *CdPlayer) String() string {
	return c.description
}
