package main

import (
	"duck/behavior/fly"
	"duck/duck"
)

func main() {
	mallard := new(duck.MallardDuck)
	rubber := new(duck.RubberDuck)
	model := new(duck.ModelDuck)
	decoy := new(duck.DecoyDuck)

	mallard.Init()
	rubber.Init()
	decoy.Init()
	model.Init()

	mallard.PerformQuack() // Quack
	rubber.PerformQuack()  // Squeak Quack
	decoy.PerformQuack()   // << Slience >>

	model.PerformFly() // I can't fly
	model.SetFlyBehavior(new(fly.FlyWithRocketPower))
	model.PerformFly() // I'm flying with a rocket
}
