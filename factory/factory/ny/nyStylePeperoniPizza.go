package ny

import (
	"factory/pizza"
)

type NYStylePepperoniPizza struct {
}

func NewNYStylePepperoniPizza() pizza.IPizza {
	var pizza NYStyleCheesePizza
	var topping []string = []string{"Grated Reggiano Cheese", "Sliced Pepperoni", "Garlic", "Onion", "Mushrooms", "Red Pepper"}
	pizza.Name = "NY Style Pepperoni Pizza"
	pizza.Dough = "Thin Crust Dough"
	pizza.Sauce = "Marinara Sauce"
	pizza.Topping = topping
	return &pizza
}
