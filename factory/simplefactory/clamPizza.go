package main

type ClamPizza struct {
}

func NewClamPizza() *Pizza {
	var topping []string = []string{"Clams", "Grated parmesan cheese"}
	return NewPizza("Clam Pizza", "Thin crust", "White garlic sauce", topping)
}
