package main

import "fmt"

type CaffeineBeverageWithHook struct {
}

func (c *CaffeineBeverageWithHook) PrepareRecipe(prepare ICaffeineBeveragePrepareWithHook) {
	c.boilWater()
	prepare.Brew()
	c.pourIntoCup()
	if prepare.CustomerWantsCondiments() {
		prepare.AddCondiments()
	}
}

func (c *CaffeineBeverageWithHook) Brew() {
	panic("implement me")
}

func (c *CaffeineBeverageWithHook) AddCondiments() {
	panic("implement me")
}

func (c *CaffeineBeverageWithHook) boilWater() {
	fmt.Println("Boiling water")
}

func (c *CaffeineBeverageWithHook) pourIntoCup() {
	fmt.Println("Pourint into cup")
}

func (c *CaffeineBeverageWithHook) CustomerWantsCondiments() bool {
	return true
}
