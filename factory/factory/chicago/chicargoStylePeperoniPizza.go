package chicago

import (
	"factory/pizza"
	"fmt"
)

type ChicagoStylePepperoniPizza struct {
	pizza.Pizza
}

func NewChicagoStylePepperoniPizza() pizza.IPizza {
	var pizza ChicagoStylePepperoniPizza
	var topping []string = []string{"Shredded Mozzarella Cheese", "Black Olives", "Spinach", "Eggplant", "Sliced Pepperoni"}
	pizza.Name = "Chicago Style Pepperoni Pizza"
	pizza.Dough = "Extra Thick Crust Dough"
	pizza.Sauce = "Plum Tomato Sauce"
	pizza.Topping = topping
	return &pizza
}

func (c *ChicagoStylePepperoniPizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
