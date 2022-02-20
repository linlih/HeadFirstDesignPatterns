package main

import (
	"factory/chicago"
	"factory/ny"
	"fmt"
)

func main() {
	nyStore := ny.NewNYPizzaStore()
	chicagoStore := chicago.NewChicagoPizzaStore()

	fmt.Println("")
	nyStore.OrderPizza(nyStore.CreatePizza("cheese"))
	chicagoStore.OrderPizza(chicagoStore.CreatePizza("cheese"))

	fmt.Println("")
	nyStore.OrderPizza(nyStore.CreatePizza("clam"))
	chicagoStore.OrderPizza(chicagoStore.CreatePizza("clam"))

	fmt.Println("")
	nyStore.OrderPizza(nyStore.CreatePizza("pepperoni"))
	chicagoStore.OrderPizza(chicagoStore.CreatePizza("pepperoni"))

	fmt.Println("")
	nyStore.OrderPizza(nyStore.CreatePizza("veggie"))
	chicagoStore.OrderPizza(chicagoStore.CreatePizza("veggie"))
}
