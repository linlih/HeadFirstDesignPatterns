package main

import "fmt"

type Projector struct {
	description string
}

func NewProjector(description string) *Projector {
	return &Projector{description: description}
}

func (p *Projector) On() {
	fmt.Println(p.description + " On")
}

func (p *Projector) Off() {
	fmt.Println(p.description + " Off")
}

func (p *Projector) WideScreenMode() {
	fmt.Println(p.description + " in wide screen mode (16x9 aspect ratio)")
}

func (p *Projector) TvMode() {
	fmt.Println(p.description + " in tv mode (4x3 aspect ratio)")
}

func (p *Projector) String() string {
	return p.description
}
