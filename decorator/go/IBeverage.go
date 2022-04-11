package main

type IBeverage interface {
	GetDescription() string
	Cost() float32
}

type Beverage struct {
	description string
}

func (b *Beverage) GetDescription() string {
	return b.description
}

func (b *Beverage) Cost() float32 {
	panic("implement me")
}
