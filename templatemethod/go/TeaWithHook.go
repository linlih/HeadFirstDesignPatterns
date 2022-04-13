package main

import "fmt"

type TeaWithHook struct {
	CaffeineBeverageWithHook
}

func NewTeaWithHook() *TeaWithHook {
	return &TeaWithHook{}
}

func (t *TeaWithHook) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *TeaWithHook) AddCondiments() {
	fmt.Println("Adding lemon")
}

func (t *TeaWithHook) CustomerWantsCondiments() bool {
	fmt.Println("Would you like milk and sugar with your coffee (y/n)? ")

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
