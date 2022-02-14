package duck

import (
	"duck/behavior/fly"
	"duck/behavior/quack"
	"fmt"
)

type DecoyDuck struct {
	Duck // go 语言中的继承方式可以使用组合模拟
}

func (d *DecoyDuck) Display() {
	fmt.Println("I'm a duck Decoy!")
}

func (d *DecoyDuck) Init() {
	d.fb = new(fly.FlyNoWay)
	d.qb = new(quack.MuteQuack)
}
