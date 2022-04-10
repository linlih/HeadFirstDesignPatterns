package simpleremote

import "fmt"

type GarageDoor struct {
}

func NewGarageDoor() *GarageDoor {
	return &GarageDoor{}
}

func (g *GarageDoor) Up() {
	fmt.Println("Garage Door is Open")
}

func (g *GarageDoor) Down() {
	fmt.Println("Garage Door is Closed")
}

func (g *GarageDoor) Stop() {
	fmt.Println("Garage Door is stopped")
}

func (g *GarageDoor) LightOn() {
	fmt.Println("Garage Light is on")
}

func (g *GarageDoor) LightOff() {
	fmt.Println("Garage Light is off")
}
