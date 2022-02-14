package main

func main() {
	mallard := new(mallardDuck)
	rubber := new(rubberDuck)
	model := new(modelDuck)
	decoy := new(DecoyDuck)

	mallard.init()
	rubber.init()
	decoy.init()
	model.init()

	mallard.performQuack() // Quack
	rubber.performQuack()  // Squeak Quack
	decoy.performQuack()   // << Slience >>

	model.performFly() // I can't fly
	model.setFlyBehavior(new(flyWithRocketPower))
	model.performFly() // I'm flying with a rocket
}
