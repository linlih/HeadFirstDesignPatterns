package pizzastore

import (
	"abstractfactory/ingredient"
	"abstractfactory/pizza"
	"fmt"
)

type NYPizzaStore struct {
	PizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{}
}

func (n *NYPizzaStore) CreatePizza(typ string) pizza.IPizza {
	var piz pizza.IPizza
	fc := ingredient.NewNYIngredientFactory()
	if typ == "cheese" {
		piz = pizza.NewCheesePizza(fc)
		piz.SetName("New York Style Cheese Pizza")
	} else if typ == "pepperoni" {
		piz = pizza.NewPepperoniPizza(fc)
		piz.SetName("New York Style Pepperoni Pizza")
	} else if typ == "clam" {
		piz = pizza.NewClamPizza(fc)
		piz.SetName("New York Style Clam Pizza")
	} else if typ == "veggie" {
		piz = pizza.NewVeggiePizza(fc)
		piz.SetName("New York Style Veggie Pizza")

	} else {
		fmt.Println("unknown type of pizza: ", typ)
	}
	return piz
}
