package main

type CheesePizza struct {
}

func NewCheesePizza() *Pizza {
	var topping []string = []string{"Fresh Mozzarella", "Parmesan"}
	return NewPizza("Cheese Pizza", "Regular Crust", "Marinara Pizza Sauce", topping)
}
