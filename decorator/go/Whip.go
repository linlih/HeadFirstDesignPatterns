package main

type Whip struct {
	beverage IBeverage
}

func NewWhip(beverage IBeverage) IBeverage {
	return &Whip{beverage: beverage}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", Whip"
}

func (w *Whip) Cost() float32 {
	return 0.10 + w.beverage.Cost()
}
