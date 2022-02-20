package main

import (
	"fmt"
)

type SimplePizzaFactory struct {
}

func NewSimplePizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}

func (s *SimplePizzaFactory) CreatePizza(typ string) *Pizza {
	if typ == "cheese" {
		return NewCheesePizza()
	} else if typ == "pepperoni" {
		return NewPepperoniPizza()
	} else if typ == "clam" {
		return NewClamPizza()
	} else if typ == "veggie" {
		return NewVeggiePizza()
	}
	fmt.Println("unknown type of pizza: ", typ)
	return nil
}
