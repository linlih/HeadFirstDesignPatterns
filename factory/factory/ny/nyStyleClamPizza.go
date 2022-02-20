package ny

import (
	"factory/pizza"
)

type NYStyleClamPizza struct {
	pizza.Pizza
}

func NewNYStyleClamPizza() pizza.IPizza {
	var pizza NYStyleCheesePizza
	var topping []string = []string{"Grated Reggiano Cheese", "Fresh Clams from long Island Sound"}
	pizza.Name = "NY Style Clam Pizza"
	pizza.Dough = "Thin Crust Dough"
	pizza.Sauce = "Marinara Sauce"
	pizza.Topping = topping
	return &pizza
}
