package main

import "fmt"

type mallardDuck struct {
	Duck
}

func (m *mallardDuck) display() {
	fmt.Println("I'm a real Mallard duck!")
}

func (m *mallardDuck) init() {
	m.setFlyBehavior(new(flyWithWings))
	m.setQuackBehavior(new(quack))
}
