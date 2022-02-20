package ny

import (
	"factory/pizza"
	"fmt"
)

type NYPizzaStore struct {
	pizza.Pizza
	pizza.PizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{}
}

func (c *NYPizzaStore) CreatePizza(typ string) pizza.IPizza {
	if typ == "cheese" {
		return NewNYStyleCheesePizza()
	} else if typ == "pepperoni" {
		return NewNYStylePepperoniPizza()
	} else if typ == "clam" {
		return NewNYStyleClamPizza()
	} else if typ == "veggie" {
		return NewNYStyleVeggiePizza()
	}
	fmt.Println("unknown type of pizza: ", typ)
	return nil
}
