package simpleremote

import "fmt"

type Light struct {
}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) On() {
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	fmt.Println("Light is off")
}
