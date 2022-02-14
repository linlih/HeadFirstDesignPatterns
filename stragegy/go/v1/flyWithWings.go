package main

import "fmt"

type flyWithWings struct {
}

func (f *flyWithWings) fly() {
	fmt.Println("I'm flying!!")
}
