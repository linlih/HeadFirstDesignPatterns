package ingredient

import (
	"abstractfactory/ingredient/cheese"
	"abstractfactory/ingredient/clams"
	"abstractfactory/ingredient/dough"
	"abstractfactory/ingredient/pepperoni"
	"abstractfactory/ingredient/sauce"
	"abstractfactory/ingredient/veggie"
)

type NYIngredientFactory struct {
}

func NewNYIngredientFactory() *NYIngredientFactory {
	return &NYIngredientFactory{}
}

func (n *NYIngredientFactory) CreateDough() string {
	crustDough := dough.ThickCrustDough{}
	return crustDough.DoughString()
}

func (n *NYIngredientFactory) CreateSauce() string {
	marinaraSauce := sauce.MarinaraSauce{}
	return marinaraSauce.SauceString()
}

func (n *NYIngredientFactory) CreateCheese() string {
	reggianoCheese := cheese.ReggianoCheese{}
	return reggianoCheese.CheeseString()
}

func (n *NYIngredientFactory) CreatePepperoni() string {
	slicedPepperoni := pepperoni.SlicedPepperoni{}
	return slicedPepperoni.PepperoniString()
}

func (n *NYIngredientFactory) CreateClams() string {
	freshClams := clams.FreshClams{}
	return freshClams.ClamsString()
}

func (n *NYIngredientFactory) CreateVeggie() []string {
	var veggies []string
	garlic := veggie.Garlic{}
	veggies = append(veggies, garlic.VeggieString())
	onion := veggie.Onion{}
	veggies = append(veggies, onion.VeggieString())
	mushroom := veggie.Mushroom{}
	veggies = append(veggies, mushroom.VeggieString())
	pepper := veggie.RedPepper{}
	veggies = append(veggies, pepper.VeggieString())
	return veggies
}
