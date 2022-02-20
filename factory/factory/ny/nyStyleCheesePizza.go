package ny

import (
	"factory/pizza"
)

type NYStyleCheesePizza struct {
	pizza.Pizza
}

func NewNYStyleCheesePizza() pizza.IPizza {
	var pizza NYStyleCheesePizza
	var topping []string = []string{"Grated Reggiano Cheese"}
	pizza.Name = "NY Style Sauce and Cheese Pizza"
	pizza.Dough = "Thin Crust Dough"
	pizza.Sauce = "Marinara Sauce"
	pizza.Topping = topping
	return &pizza
}
