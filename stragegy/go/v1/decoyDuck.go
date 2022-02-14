package main

import "fmt"

type DecoyDuck struct {
	Duck // go 语言中的继承方式可以使用组合模拟
}

func (d *DecoyDuck) display() {
	fmt.Println("I'm a duck Decoy!")
}

func (d *DecoyDuck) init() {
	d.fb = new(flyNoWay)
	d.qb = new(muteQuack)
}
