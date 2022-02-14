package duck

import (
	"duck/behavior/fly"
	"duck/behavior/quack"
	"fmt"
)

type MallardDuck struct {
	Duck
}

func (m *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck!")
}

func (m *MallardDuck) Init() {
	m.SetFlyBehavior(new(fly.FlyWithWings))
	m.SetQuackBehavior(new(quack.Quack))
}
