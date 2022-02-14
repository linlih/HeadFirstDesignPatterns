package duck

import (
	"duck/behavior/fly"
	"duck/behavior/quack"
	"fmt"
)

type RedHeadDuck struct {
	Duck
}

func (r *RedHeadDuck) Display() {
	fmt.Println("I'm a read Red Head duck!")
}

func (r *RedHeadDuck) Init() {
	r.SetFlyBehavior(new(fly.FlyWithWings))
	r.SetQuackBehavior(new(quack.Quack))
}
