package main

import "fmt"

type flyWithRocketPower struct {
}

func (f *flyWithRocketPower) fly() {
	fmt.Println("I'm flying with a rocket")
}
