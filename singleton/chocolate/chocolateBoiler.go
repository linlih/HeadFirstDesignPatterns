package main

import (
	"fmt"
	"sync"
)

type chocolateBoiler struct {
	empty  bool
	boiled bool
}

var once sync.Once

var chocolateBoilerInstance *chocolateBoiler

func getChocolateBoilerInstance() *chocolateBoiler {
	if chocolateBoilerInstance == nil {
		once.Do(func() {
			fmt.Println("Create chocolate Boiler now.")
			chocolateBoilerInstance = &chocolateBoiler{empty: true, boiled: false}
		})
	} else {
		fmt.Println("Chocolate Boiler already created.")
	}
	return chocolateBoilerInstance
}

func (c *chocolateBoiler) Fill() {
	if c.isEmpty() {
		c.empty = false
		c.boiled = false
		// fill the boiler with a milk/chocolate mixture
	}
}

func (c *chocolateBoiler) Drain() {
	if !c.isEmpty() && !c.isBoiled() {
		c.empty = true
	}
}

func (c *chocolateBoiler) Boil() {
	if !c.isEmpty() && !c.isBoiled() {
		c.boiled = true
	}
}

func (c *chocolateBoiler) isEmpty() bool {
	return c.empty
}

func (c *chocolateBoiler) isBoiled() bool {
	return c.boiled
}
