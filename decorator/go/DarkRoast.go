package main

type DarkRoast struct {
	Beverage
}

func NewDarkRoast() IBeverage {
	d := &DarkRoast{}
	d.description = "Dark Roast Coffee"
	return d
}

func (d *DarkRoast) Cost() float32 {
	return 0.99
}
