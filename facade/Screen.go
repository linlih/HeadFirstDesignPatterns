package main

import "fmt"

type Screen struct {
	description string
}

func NewScreen(description string) *Screen {
	return &Screen{description: description}
}

func (s *Screen) Up() {
	fmt.Println(s.description + " Up")
}

func (s *Screen) Down() {
	fmt.Println(s.description + " Down")
}

func (s *Screen) String() string {
	return s.description
}
