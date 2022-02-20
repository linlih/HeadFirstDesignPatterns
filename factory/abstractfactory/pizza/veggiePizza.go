package pizza

import (
	"abstractfactory/ingredient"
	"fmt"
)

type VeggiePizza struct {
	Pizza
	fc ingredient.IIngredientFactory
}

func NewVeggiePizza(factory ingredient.IIngredientFactory) IPizza {
	return &VeggiePizza{fc: factory}
}

func (v *VeggiePizza) Prepare() {
	fmt.Println("Preparing " + v.GetName())
	v.Dough = v.fc.CreateDough()
	v.Sauce = v.fc.CreateSauce()
	v.Cheese = v.fc.CreateCheese()
	v.Clams = v.fc.CreateClams()
}
