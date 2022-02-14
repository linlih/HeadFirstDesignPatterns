package main

import "fmt"

type redHeadDuck struct {
	Duck
}

func (r *redHeadDuck) display() {
	fmt.Println("I'm a read Red Head duck!")
}

func (r *redHeadDuck) init() {
	r.setFlyBehavior(new(flyWithWings))
	r.setQuackBehavior(new(quack))
}
