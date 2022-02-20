package main

import (
	"abstractfactory/pizzastore"
	"fmt"
)

func main() {
	nyStore := pizzastore.NewNYPizzaStore()
	chicagoStore := pizzastore.NewChicargoPizzaStore()

	pizza := nyStore.OrderPizza(nyStore.CreatePizza("cheese"))
	fmt.Println("Ethan ordered a " + pizza.String() + "\n")

	pizza = chicagoStore.OrderPizza(chicagoStore.CreatePizza("cheese"))
	fmt.Println("Joel ordered a " + pizza.String() + "\n")

	pizza = nyStore.OrderPizza(nyStore.CreatePizza("clam"))
	fmt.Println("Ethan ordered a " + pizza.String() + "\n")

	pizza = chicagoStore.OrderPizza(nyStore.CreatePizza("clam"))
	fmt.Println("Joel ordered a " + pizza.String() + "\n")

	pizza = nyStore.OrderPizza(chicagoStore.CreatePizza("pepperoni"))
	fmt.Println("Ethan ordered a " + pizza.String() + "\n")

	pizza = chicagoStore.OrderPizza(nyStore.CreatePizza("pepperoni"))
	fmt.Println("Joel ordered a " + pizza.String() + "\n")

	pizza = nyStore.OrderPizza(nyStore.CreatePizza("veggie"))
	fmt.Println("Ethan ordered a " + pizza.String() + "\n")

	pizza = chicagoStore.OrderPizza(chicagoStore.CreatePizza("veggie"))
	fmt.Println("Joel ordered a " + pizza.String() + "\n")
}
