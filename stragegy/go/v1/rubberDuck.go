package main

import "fmt"

type rubberDuck struct {
	Duck
}

func (r *rubberDuck) display() {
	fmt.Println("I'm a read Rubber duckie!")
}

func (r *rubberDuck) init() {
	r.setFlyBehavior(new(flyWithWings))
	r.setQuackBehavior(new(squeakQuack))
}
