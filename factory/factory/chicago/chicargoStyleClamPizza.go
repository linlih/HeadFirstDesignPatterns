package chicago

import (
	"factory/pizza"
	"fmt"
)

type ChicagoStyleClamPizza struct {
	pizza.Pizza
}

func NewChicagoStyleClamPizza() pizza.IPizza {
	var pizza ChicagoStyleClamPizza
	var topping []string = []string{"Shredded Mozzarella Cheese", "Frozen Clams from Chesapeake Bay"}
	pizza.Name = "Chicago Style Clam Pizza"
	pizza.Dough = "Extra Thick Crust Dough"
	pizza.Sauce = "Plum Tomato Sauce"
	pizza.Topping = topping
	return &pizza
}

func (c *ChicagoStyleClamPizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
