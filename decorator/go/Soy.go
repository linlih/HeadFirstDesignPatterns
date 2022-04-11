package main

type Soy struct {
	beverage IBeverage
}

func NewSoy(beverage IBeverage) IBeverage {
	return &Soy{beverage: beverage}
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", Soy"
}

func (s *Soy) Cost() float32 {
	return 0.15 + s.beverage.Cost()
}
