package adapter

import (
	"adapter/duck"
	"math/rand"
)

type DuckAdapter struct {
	duck duck.IDuck
}

func NewDuckAdapter(duck duck.IDuck) *DuckAdapter {
	return &DuckAdapter{duck: duck}
}

func (d *DuckAdapter) Gobble() {
	d.duck.Quack()
}

func (d *DuckAdapter) Fly() {
	for rand.Intn(5) == 0 {
		d.duck.Fly()
	}
}
