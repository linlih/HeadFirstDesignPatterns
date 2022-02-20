package ny

import (
	"factory/pizza"
)

type NYStyleVeggiePizza struct {
}

func NewNYStyleVeggiePizza() pizza.IPizza {
	var pizza NYStyleCheesePizza
	var topping []string = []string{"Grated Reggiano Cheese", "Garlic", "Onion", "Mushrooms", "Red Pepper"}
	pizza.Name = "NY Style Veggie Pizza"
	pizza.Dough = "Thin Crust Dough"
	pizza.Sauce = "Marinara Sauce"
	pizza.Topping = topping
	return &pizza
}
