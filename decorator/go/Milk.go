package main

type Milk struct {
	beverage IBeverage
}

func NewMilk(beverage IBeverage) IBeverage {
	return &Milk{beverage: beverage}
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float32 {
	return 0.10 + m.beverage.Cost()
}
