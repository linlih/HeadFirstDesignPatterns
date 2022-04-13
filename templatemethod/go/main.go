package main

func main() {
	tea := NewTea()
	coffee := NewCoffee()
	tea.PrepareRecipe(tea)
	coffee.PrepareRecipe(coffee)

	teahook := NewTeaWithHook()
	coffeehook := NewCoffeeWithHook()
	teahook.PrepareRecipe(teahook)
	coffeehook.PrepareRecipe(coffeehook)
}
