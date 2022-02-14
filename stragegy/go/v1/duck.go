package main

import "fmt"

type Duck struct {
	fb iFlyBehavior
	qb iQuackBehavior
}

func (d *Duck) setFlyBehavior(fb iFlyBehavior) {
	d.fb = fb
}

func (d *Duck) setQuackBehavior(qb iQuackBehavior) {
	d.qb = qb
}

func (d *Duck) display() {
	fmt.Println("duck display")
}

func (d *Duck) performFly() {
	d.fb.fly()
}

func (d *Duck) performQuack() {
	d.qb.quack()
}

func (d *Duck) swin() {
	fmt.Println("All ducks float, even decoys!")
}
