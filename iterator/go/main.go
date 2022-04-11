package main

func main() {
	pancakeHouseMenu := NewPancakeHouseMenu()
	dinerMenu := NewDinerMenu()

	waitress := NewWaitress(pancakeHouseMenu, dinerMenu)
	waitress.PrintMenu()

	waitress.PrintVegetarianMenu()
}
