package main

import "fmt"

type CoffeeWithHook struct {
	CaffeineBeverageWithHook
}

func NewCoffeeWithHook() *CoffeeWithHook {
	return &CoffeeWithHook{}
}

func (c *CoffeeWithHook) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *CoffeeWithHook) AddCondiments() {
	fmt.Println("Add Sugar and Milk")
}

func (c *CoffeeWithHook) CustomerWantsCondiments() bool {
	fmt.Println("Would you like lemon with your tea (y/n)? ")

	var answer string

	for {
		fmt.Scanln(&answer)
		if answer == "y" {
			return true
		} else if answer == "n" {
			return false
		} else {
			fmt.Println("Wrong input, please input y or n.")
		}
	}
}
