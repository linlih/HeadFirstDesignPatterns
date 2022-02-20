package chicago

import (
	"factory/pizza"
	"fmt"
)

type ChicagoStyleVeggiePizza struct {
	pizza.Pizza
}

func NewChicagoStyleVeggiePizza() pizza.IPizza {
	var pizza ChicagoStyleVeggiePizza
	var topping []string = []string{"Shredded Mozzarella Cheese", "Black Olives", "Spinach", "Eggplant"}
	pizza.Name = "Chicago Deep Dish Veggie Pizza"
	pizza.Dough = "Extra Thick Crust Dough"
	pizza.Sauce = "Plum Tomato Sauce"
	pizza.Topping = topping
	return &pizza
}

func (c *ChicagoStyleVeggiePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
