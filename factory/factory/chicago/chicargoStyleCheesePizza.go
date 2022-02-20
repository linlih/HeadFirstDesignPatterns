package chicago

import (
	"factory/pizza"
	"fmt"
)

type ChicagoStyleCheesePizza struct {
	pizza.Pizza
}

func NewChicagoStyleCheesePizza() pizza.IPizza {
	var pizza ChicagoStyleCheesePizza
	var topping []string = []string{"Shredded Mozzarella Cheese"}
	pizza.Name = "Chicago Style Deep Dish Cheese Pizza"
	pizza.Dough = "Extra Thick Crust Dough"
	pizza.Sauce = "Plum Tomato Sauce"
	pizza.Topping = topping
	return &pizza
}

func (c *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
