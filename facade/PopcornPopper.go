package main

import "fmt"

type PopcornPopper struct {
	description string
}

func NewPopcornPopper(description string) *PopcornPopper {
	return &PopcornPopper{description: description}
}

func (p *PopcornPopper) On() {
	fmt.Println(p.description + " On")
}

func (p *PopcornPopper) Off() {
	fmt.Println(p.description + " Off")
}

func (p *PopcornPopper) Pop() {
	fmt.Println(p.description + " popping popcorn!")
}

func (p *PopcornPopper) String() string {
	return p.description
}
