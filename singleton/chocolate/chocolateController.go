package main

func main() {
	boiler := getChocolateBoilerInstance()

	boiler.Fill()
	boiler.Boil()
	boiler.Drain()

	_ = getChocolateBoilerInstance()
}
