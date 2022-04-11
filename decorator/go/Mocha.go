package main

type Mocha struct {
	beverage IBeverage
}

func NewMocha(beverage IBeverage) IBeverage {
	return &Mocha{beverage: beverage}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float32 {
	return 0.20 + m.beverage.Cost()
}
