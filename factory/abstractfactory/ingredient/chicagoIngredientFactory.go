package ingredient

import (
	"abstractfactory/ingredient/cheese"
	"abstractfactory/ingredient/clams"
	"abstractfactory/ingredient/dough"
	"abstractfactory/ingredient/pepperoni"
	"abstractfactory/ingredient/sauce"
	"abstractfactory/ingredient/veggie"
)

type ChicagoIngredientFactory struct {
}

func NewChicagoIngredientFactory() *ChicagoIngredientFactory {
	return &ChicagoIngredientFactory{}
}

func (c *ChicagoIngredientFactory) CreateDough() string {
	crustDough := dough.ThickCrustDough{}
	return crustDough.DoughString()
}

func (c *ChicagoIngredientFactory) CreateSauce() string {
	tomatoSauce := sauce.PlumTomatoSauce{}
	return tomatoSauce.SauceString()
}

func (c *ChicagoIngredientFactory) CreateCheese() string {
	mozzarellaCheese := cheese.MozzarellaCheese{}
	return mozzarellaCheese.CheeseString()
}

func (c *ChicagoIngredientFactory) CreatePepperoni() string {
	slicedPepperoni := pepperoni.SlicedPepperoni{}
	return slicedPepperoni.PepperoniString()
}

func (c *ChicagoIngredientFactory) CreateClams() string {
	frozenClams := clams.FrozenClams{}
	return frozenClams.ClamsString()
}

func (c *ChicagoIngredientFactory) CreateVeggie() []string {
	var veggies []string
	olives := veggie.BlackOlives{}
	veggies = append(veggies, olives.VeggieString())
	spinach := veggie.Spinach{}
	veggies = append(veggies, spinach.VeggieString())
	eggplant := veggie.Eggplant{}
	veggies = append(veggies, eggplant.VeggieString())
	return veggies
}
