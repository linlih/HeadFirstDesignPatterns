package main

type VeggiePizza struct {
}

func NewVeggiePizza() *Pizza {
	var topping []string = []string{"Shredded mozzarella", "Grated parmesan", "Diced onion", "Sliced mushrooms", "Sliced red pepper", "Sliced black olives"}
	return NewPizza("Veggie Pizza", "Crust", "Marinara sauce", topping)
}
