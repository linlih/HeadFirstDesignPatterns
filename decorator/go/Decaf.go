package main

type Decaf struct {
	Beverage
}

func NewDecaf() IBeverage {
	d := &DarkRoast{}
	d.description = "Decaf Coffee"
	return d
}

func (d *Decaf) Cost() float32 {
	return 1.05
}
