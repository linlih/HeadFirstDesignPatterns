package duck

import (
	"duck/behavior/fly"
	"duck/behavior/quack"
	"fmt"
)

type Duck struct {
	// strategy 模式的核心
	fb fly.IFlyBehavior
	qb quack.IQuackBehavior
}

func (d *Duck) SetFlyBehavior(fb fly.IFlyBehavior) {
	d.fb = fb
}

func (d *Duck) SetQuackBehavior(qb quack.IQuackBehavior) {
	d.qb = qb
}

func (d *Duck) Display() {
	fmt.Println("duck display")
}

func (d *Duck) PerformFly() {
	d.fb.Fly()
}

func (d *Duck) PerformQuack() {
	d.qb.Quack()
}

func (d *Duck) Swin() {
	fmt.Println("All ducks float, even decoys!")
}
