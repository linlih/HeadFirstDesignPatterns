package duck

import (
	"duck/behavior/fly"
	"duck/behavior/quack"
	"fmt"
)

type RubberDuck struct {
	Duck
}

func (r *RubberDuck) Display() {
	fmt.Println("I'm a read Rubber duckie!")
}

func (r *RubberDuck) Init() {
	r.SetFlyBehavior(new(fly.FlyWithWings))
	r.SetQuackBehavior(new(quack.SqueakQuack))
}
