package main

type Espresso struct {
	Beverage
}

func NewEspresso() IBeverage {
	d := &Espresso{}
	d.description = "Espresso"
	return d
}

func (d *Espresso) Cost() float32 {
	return 1.99
}
