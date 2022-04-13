package main

import "fmt"

type CaffeineBeverage struct {
}

func (c *CaffeineBeverage) PrepareRecipe(prepare ICaffeineBeveragePrepare) {
	c.boilWater()
	prepare.Brew()
	c.pourIntoCup()
	prepare.AddCondiments()
}

func (c *CaffeineBeverage) Brew() {
	panic("implement me")
}

func (c *CaffeineBeverage) AddCondiments() {
	panic("implement me")
}

func (c *CaffeineBeverage) boilWater() {
	fmt.Println("Boiling water")
}

func (c *CaffeineBeverage) pourIntoCup() {
	fmt.Println("Pourint into cup")
}
