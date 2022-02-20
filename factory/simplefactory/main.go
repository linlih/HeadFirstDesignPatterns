package main

import "fmt"

func main() {
	factory := NewSimplePizzaFactory()
	store := NewPizzaStore(factory)

	pizza := store.orderPizza("cheese")
	fmt.Println("We ordered a", pizza.GetName())

	pizza = store.orderPizza("pepperoni")
	fmt.Println("We ordered a", pizza.GetName())

	pizza = store.orderPizza("clam")
	fmt.Println("We ordered a", pizza.GetName())

	pizza = store.orderPizza("veggie")
	fmt.Println("We ordered a", pizza.GetName())
}
