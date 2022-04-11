package main

type HouseBlend struct {
	Beverage
}

func NewHouseBlend() IBeverage {
	d := &HouseBlend{}
	d.description = "Dark Roast Coffee"
	return d
}

func (d *HouseBlend) Cost() float32 {
	return 0.99
}
