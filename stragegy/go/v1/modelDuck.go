package main

import "fmt"

type modelDuck struct {
	Duck
}

func (m *modelDuck) display() {
	fmt.Println("I'm a model duck!")
}

func (m *modelDuck) init() {
	m.setFlyBehavior(new(flyNoWay))
	m.setQuackBehavior(new(quack))
}
