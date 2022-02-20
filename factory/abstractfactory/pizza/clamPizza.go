package pizza

import (
	"abstractfactory/ingredient"
	"fmt"
)

type ClamPizza struct {
	Pizza
	fc ingredient.IIngredientFactory
}

func NewClamPizza(factory ingredient.IIngredientFactory) IPizza {
	return &ClamPizza{fc: factory}
}

func (c *ClamPizza) Prepare() {
	fmt.Println("Preparing " + c.GetName())
	c.Dough = c.fc.CreateDough()
	c.Sauce = c.fc.CreateSauce()
	c.Cheese = c.fc.CreateCheese()
	c.Clams = c.fc.CreateClams()
}
