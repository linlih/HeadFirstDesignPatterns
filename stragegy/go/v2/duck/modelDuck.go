package duck

import (
	"duck/behavior/fly"
	"duck/behavior/quack"
	"fmt"
)

type ModelDuck struct {
	Duck
}

func (m *ModelDuck) Display() {
	fmt.Println("I'm a model duck!")
}

func (m *ModelDuck) Init() {
	m.SetFlyBehavior(new(fly.FlyNoWay))
	m.SetQuackBehavior(new(quack.Quack))
}
