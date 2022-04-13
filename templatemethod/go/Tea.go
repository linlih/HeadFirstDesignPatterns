package main

import "fmt"

type Tea struct {
	CaffeineBeverage
}

func NewTea() ICaffeineBeveragePrepare {
	return &Tea{}
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}
