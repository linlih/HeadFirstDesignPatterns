package chicago

import (
	"factory/pizza"
	"fmt"
)

type ChicagoPizzaStore struct {
	pizza.Pizza
	pizza.PizzaStore
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	return &ChicagoPizzaStore{}
}

func (c *ChicagoPizzaStore) CreatePizza(typ string) pizza.IPizza {
	if typ == "cheese" {
		return NewChicagoStyleCheesePizza()
	} else if typ == "pepperoni" {
		return NewChicagoStylePepperoniPizza()
	} else if typ == "clam" {
		return NewChicagoStyleClamPizza()
	} else if typ == "veggie" {
		return NewChicagoStyleVeggiePizza()
	}
	fmt.Println("unknown type of pizza: ", typ)
	return nil
}
