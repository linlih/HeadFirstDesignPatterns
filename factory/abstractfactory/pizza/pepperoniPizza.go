package pizza

import (
	"abstractfactory/ingredient"
	"fmt"
)

type PepperoniPizza struct {
	Pizza
	fc ingredient.IIngredientFactory
}

func NewPepperoniPizza(factory ingredient.IIngredientFactory) IPizza {
	return &PepperoniPizza{fc: factory}
}

func (p *PepperoniPizza) Prepare() {
	fmt.Println("Preparing " + p.GetName())
	p.Dough = p.fc.CreateDough()
	p.Sauce = p.fc.CreateSauce()
	p.Cheese = p.fc.CreateCheese()
}
