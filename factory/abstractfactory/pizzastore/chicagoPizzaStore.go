package pizzastore

import (
	"abstractfactory/ingredient"
	"abstractfactory/pizza"
	"fmt"
)

type ChicargoPizzaStore struct {
	PizzaStore
}

func NewChicargoPizzaStore() *ChicargoPizzaStore {
	return &ChicargoPizzaStore{}
}

func (c *ChicargoPizzaStore) CreatePizza(typ string) pizza.IPizza {
	var piz pizza.IPizza
	fc := ingredient.NewChicagoIngredientFactory()
	if typ == "cheese" {
		piz = pizza.NewCheesePizza(fc)
		piz.SetName("Chicago Style Cheese Pizza")
	} else if typ == "pepperoni" {
		piz = pizza.NewPepperoniPizza(fc)
		piz.SetName("Chicago Style Pepperoni Pizza")
	} else if typ == "clam" {
		piz = pizza.NewClamPizza(fc)
		piz.SetName("Chicago Style Clam Pizza")
	} else if typ == "veggie" {
		piz = pizza.NewVeggiePizza(fc)
		piz.SetName("Chicago Style Veggie Pizza")
	}
	fmt.Println("unknown type of pizza: ", typ)
	return piz
}
