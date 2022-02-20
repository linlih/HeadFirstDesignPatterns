package main

type PepperoniPizza struct {
}

func NewPepperoniPizza() *Pizza {
	var topping []string = []string{"Sliced Pepperoni", "Sliced Onion", "Grated parmesan cheese"}
	return NewPizza("Pepperoni Pizza", "Crust", "Marinara sauce", topping)
}
