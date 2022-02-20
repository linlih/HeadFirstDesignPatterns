package pizza

import (
	"abstractfactory/ingredient"
	"fmt"
)

type CheesePizza struct {
	Pizza
	fc ingredient.IIngredientFactory
}

func NewCheesePizza(factory ingredient.IIngredientFactory) IPizza {
	return &CheesePizza{fc: factory}
}

func (c *CheesePizza) Prepare() {
	fmt.Println("Preparing " + c.GetName())
	c.Dough = c.fc.CreateDough()
	c.Sauce = c.fc.CreateSauce()
	c.Cheese = c.fc.CreateCheese()
}
